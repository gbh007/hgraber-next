package rebuilder

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

var (
	errForbiddenMerge      = errors.New("merge with book forbidden")
	errEmptyPagesOnRebuild = errors.New("empty pages on rebuild")
	errMissingSourcePage   = errors.New("missing source page")
)

func (uc *UseCase) RebuildBook(ctx context.Context, request entities.RebuildBookRequest) (_ uuid.UUID, returnErr error) {
	if len(request.SelectedPages) == 0 {
		return uuid.Nil, errEmptyPagesOnRebuild
	}

	targetBookID := request.MergeWithBook
	isNewBook := request.MergeWithBook == uuid.Nil

	var (
		err              error
		bookToMerge      entities.Book
		attributeToMerge map[string][]string
		existsPageHashes map[entities.FileHash]struct{}
		existsDeadHashes map[entities.FileHash]struct{}
	)

	if !isNewBook {
		bookToMerge, err = uc.storage.GetBook(ctx, request.MergeWithBook)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: get book to merge: %w", err)
		}

		if !bookToMerge.IsRebuild {
			return uuid.Nil, fmt.Errorf("%w: not rebuilded book", errForbiddenMerge)
		}

		if bookToMerge.Deleted {
			return uuid.Nil, fmt.Errorf("%w: deleted book", errForbiddenMerge)
		}

		attributeToMerge, err = uc.storage.BookOriginAttributes(ctx, request.MergeWithBook)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: get attributes to merge: %w", err)
		}

		if request.OnlyUniquePages {
			pages, err := uc.storage.BookPagesWithHash(ctx, request.MergeWithBook)
			if err != nil {
				return uuid.Nil, fmt.Errorf("storage: get page to unique: %w", err)
			}

			existsPageHashes = make(map[entities.FileHash]struct{}, len(pages)+len(request.SelectedPages))

			for _, page := range pages {
				existsPageHashes[page.FileHash] = struct{}{}
			}
		}
	} else {
		targetBookID = uuid.Must(uuid.NewV7())
		bookToMerge = entities.Book{
			ID:         targetBookID,
			Name:       request.OldBook.Book.Name,
			OriginURL:  request.OldBook.Book.OriginURL,
			PageCount:  0,
			IsRebuild:  true,
			CreateAt:   time.Now().UTC(),
			Verified:   true,
			VerifiedAt: time.Now().UTC(),
		}

		if request.OnlyUniquePages {
			existsPageHashes = make(map[entities.FileHash]struct{}, len(request.SelectedPages))
		}
	}

	sourcePages, err := uc.storage.BookPagesWithHash(ctx, request.OldBook.Book.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("storage: get source pages: %w", err)
	}

	sourcePagesMap := make(map[int]entities.PageWithHash)

	for _, page := range sourcePages {
		sourcePagesMap[page.PageNumber] = page
	}

	if request.ExcludeDeadHashPages {
		md5Sums := make([]string, 0, len(sourcePages))

		for _, page := range sourcePages {
			md5Sums = append(md5Sums, page.Md5Sum)
		}

		deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: get dead hashes: %w", err)
		}

		existsDeadHashes = make(map[entities.FileHash]struct{}, len(deadHashes))

		for _, hash := range deadHashes {
			existsDeadHashes[hash.FileHash] = struct{}{}
		}
	}

	request.SelectedPages = slices.Compact(request.SelectedPages)
	slices.Sort(request.SelectedPages)

	pagesRemap := make(map[int]int, len(request.SelectedPages))
	newPages := make([]entities.Page, 0, len(request.SelectedPages))

	newPageNumberCounter := 0

	for _, oldPageNumber := range request.SelectedPages {
		sourcePage, ok := sourcePagesMap[oldPageNumber]
		if !ok {
			return uuid.Nil, fmt.Errorf("%w (%d)", errMissingSourcePage, oldPageNumber)
		}

		if request.ExcludeDeadHashPages {
			if _, ok := existsDeadHashes[sourcePage.FileHash]; ok {
				continue
			}
		}

		if request.OnlyUniquePages {
			if _, ok := existsPageHashes[sourcePage.FileHash]; ok {
				continue
			}

			existsPageHashes[sourcePage.FileHash] = struct{}{}
		}

		newPageNumberCounter++
		newPageNumber := bookToMerge.PageCount + newPageNumberCounter
		pagesRemap[oldPageNumber] = newPageNumber

		sourcePage.BookID = targetBookID
		sourcePage.PageNumber = newPageNumber
		sourcePage.CreateAt = time.Now()

		newPages = append(newPages, sourcePage.Page)
	}

	bookToMerge.PageCount += newPageNumberCounter
	bookToMerge.AttributesParsed = true

	if isNewBook && len(newPages) == 0 {
		return uuid.Nil, fmt.Errorf("%w: after deduplicate for new book", errEmptyPagesOnRebuild)
	}

	if isNewBook {
		err = uc.storage.NewBook(ctx, bookToMerge)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: create book: %w", err)
		}

		defer func() { // Удаляем книгу, если не получилось ее создать полностью
			if returnErr != nil {
				deleteErr := uc.storage.DeleteBook(ctx, targetBookID)
				if deleteErr != nil {
					uc.logger.ErrorContext(
						ctx, "delete new book after unsuccess rebuild",
						slog.Any("error", deleteErr),
						slog.String("book_id", targetBookID.String()),
					)
				}
			}
		}()
	} else {
		err = uc.storage.UpdateBook(ctx, bookToMerge)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: create book: %w", err)
		}
	}

	newLabels := make([]entities.BookLabel, 0, len(request.OldBook.Labels))

	for _, label := range request.OldBook.Labels {
		_, ok := pagesRemap[label.PageNumber]
		if !ok && label.PageNumber != 0 { // Отсекаем данные которые не были замаплены или не привязаны к книге.
			continue
		}

		label.BookID = targetBookID
		label.PageNumber = pagesRemap[label.PageNumber]

		newLabels = append(newLabels, label)
	}

	newAttributes := entities.MergeAttributeMap(attributeToMerge, request.OldBook.Attributes)

	if len(newPages) > 0 {
		err = uc.storage.NewBookPages(ctx, newPages)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: new pages: %w", err)
		}
	}

	if len(newAttributes) > 0 {
		err = uc.storage.UpdateOriginAttributes(ctx, targetBookID, newAttributes)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: set origin attributes: %w", err)
		}

		err = uc.storage.UpdateAttributes(ctx, targetBookID, newAttributes)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: set attributes: %w", err)
		}
	}

	if len(newLabels) > 0 {
		err = uc.storage.SetLabels(ctx, newLabels)
		if err != nil {
			return uuid.Nil, fmt.Errorf("storage: set labels: %w", err)
		}
	}

	return targetBookID, nil
}

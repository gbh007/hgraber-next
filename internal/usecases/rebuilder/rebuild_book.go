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

	isNewBook := request.MergeWithBook == uuid.Nil

	bookToMerge, attributeToMerge, targetPageHashes, err := uc.rebuildBookGetTarget(ctx, request)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: get target: %w", err)
	}

	sourcePagesMap, forbiddenHashes, err := uc.rebuildBookPrepareResources(
		ctx,
		request.Flags,
		request.OldBook.Book.ID,
		targetPageHashes,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: prepare resources: %w", err)
	}

	pagesRemap, newPages, err := uc.rebuildBookPages(
		ctx,
		request.Flags,
		slices.Clone(request.SelectedPages),
		&bookToMerge,
		sourcePagesMap,
		forbiddenHashes,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: pages: %w", err)
	}

	newLabels := make([]entities.BookLabel, 0, len(request.OldBook.Labels))

	for _, label := range request.OldBook.Labels {
		_, ok := pagesRemap[label.PageNumber]
		if !ok && label.PageNumber != 0 { // Отсекаем данные которые не были замаплены или не привязаны к книге.
			continue
		}

		label.BookID = bookToMerge.ID
		label.PageNumber = pagesRemap[label.PageNumber]

		newLabels = append(newLabels, label)
	}

	bookToMerge.AttributesParsed = true
	newAttributes := entities.MergeAttributeMap(attributeToMerge, request.OldBook.Attributes)

	if isNewBook && len(newPages) == 0 {
		return uuid.Nil, fmt.Errorf("%w: after deduplicate for new book", errEmptyPagesOnRebuild)
	}

	err = uc.rebuildBookSave(
		ctx,
		isNewBook,
		bookToMerge,
		newAttributes,
		newPages, newLabels,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: save: %w", err)
	}

	return bookToMerge.ID, nil
}

func (uc *UseCase) rebuildBookGetTarget(ctx context.Context, request entities.RebuildBookRequest) (
	entities.Book,
	map[string][]string,
	map[entities.FileHash]struct{},
	error,
) {
	isNewBook := request.MergeWithBook == uuid.Nil

	var (
		err              error
		bookToMerge      entities.Book
		attributeToMerge map[string][]string
		targetPageHashes map[entities.FileHash]struct{}
	)

	if isNewBook {
		bookToMerge = entities.Book{
			ID:         uuid.Must(uuid.NewV7()),
			Name:       request.OldBook.Book.Name,
			OriginURL:  request.OldBook.Book.OriginURL,
			PageCount:  0,
			IsRebuild:  true,
			CreateAt:   time.Now().UTC(),
			Verified:   true,
			VerifiedAt: time.Now().UTC(),
		}

		return bookToMerge, nil, nil, nil
	}

	bookToMerge, err = uc.storage.GetBook(ctx, request.MergeWithBook)
	if err != nil {
		return entities.Book{}, nil, nil, fmt.Errorf("storage: get book to merge: %w", err)
	}

	if !bookToMerge.IsRebuild {
		return entities.Book{}, nil, nil, fmt.Errorf("%w: not rebuilded book", errForbiddenMerge)
	}

	if bookToMerge.Deleted {
		return entities.Book{}, nil, nil, fmt.Errorf("%w: deleted book", errForbiddenMerge)
	}

	attributeToMerge, err = uc.storage.BookOriginAttributes(ctx, request.MergeWithBook)
	if err != nil {
		return entities.Book{}, nil, nil, fmt.Errorf("storage: get attributes to merge: %w", err)
	}

	if request.Flags.OnlyUniquePages {
		pages, err := uc.storage.BookPagesWithHash(ctx, request.MergeWithBook)
		if err != nil {
			return entities.Book{}, nil, nil, fmt.Errorf("storage: get page to unique: %w", err)
		}

		targetPageHashes = make(map[entities.FileHash]struct{}, len(pages)+len(request.SelectedPages))

		for _, page := range pages {
			targetPageHashes[page.FileHash] = struct{}{}
		}
	}

	return bookToMerge, attributeToMerge, targetPageHashes, nil
}

func (uc *UseCase) rebuildBookPrepareResources(
	ctx context.Context,
	flags entities.RebuildBookRequestFlags,
	oldBookID uuid.UUID,
	targetPageHashes map[entities.FileHash]struct{},
) (
	map[int]entities.PageWithHash,
	map[entities.FileHash]struct{},
	error,
) {
	sourcePages, err := uc.storage.BookPagesWithHash(ctx, oldBookID)
	if err != nil {
		return nil, nil, fmt.Errorf("storage: get source pages: %w", err)
	}

	sourcePagesMap := make(map[int]entities.PageWithHash, len(sourcePages))
	sourcePagesHashes := make(map[entities.FileHash]struct{}, len(sourcePages))
	sourcePagesMD5Sums := make([]string, 0, len(sourcePages))

	for _, page := range sourcePages {
		sourcePagesMap[page.PageNumber] = page
		sourcePagesMD5Sums = append(sourcePagesMD5Sums, page.Md5Sum)
		sourcePagesHashes[page.FileHash] = struct{}{}
	}

	forbiddenHashes := make(map[entities.FileHash]struct{}, len(targetPageHashes)+len(sourcePagesHashes))

	if flags.ExcludeDeadHashPages {
		deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, sourcePagesMD5Sums)
		if err != nil {
			return nil, nil, fmt.Errorf("storage: get dead hashes: %w", err)
		}

		for _, hash := range deadHashes {
			forbiddenHashes[hash.FileHash] = struct{}{}
		}
	}

	if flags.Only1CopyPages {
		pages, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, sourcePagesMD5Sums)
		if err != nil {
			return nil, nil, fmt.Errorf("storage: get pages by source hashes: %w", err)
		}

		for _, page := range pages {
			// Отсекаем не совпадения из-за неполного ограничения в БД
			if _, ok := sourcePagesHashes[page.FileHash]; !ok {
				continue
			}

			// Это текущая книга, ее исключать не надо
			if page.BookID == oldBookID {
				continue
			}

			forbiddenHashes[page.FileHash] = struct{}{}
		}
	}

	if flags.OnlyUniquePages || flags.Only1CopyPages {
		for hash := range targetPageHashes {
			forbiddenHashes[hash] = struct{}{}
		}
	}

	return sourcePagesMap, forbiddenHashes, nil
}

func (uc *UseCase) rebuildBookPages(
	_ context.Context,
	flags entities.RebuildBookRequestFlags,
	selectedPages []int,
	bookToMerge *entities.Book,
	sourcePagesMap map[int]entities.PageWithHash,
	forbiddenHashes map[entities.FileHash]struct{},
) (map[int]int, []entities.Page, error) {
	selectedPages = slices.Compact(selectedPages)
	slices.Sort(selectedPages)

	existsPageHashes := make(map[entities.FileHash]struct{}, len(selectedPages))

	pagesRemap := make(map[int]int, len(selectedPages))
	newPages := make([]entities.Page, 0, len(selectedPages))

	newPageNumberCounter := 0

	for _, oldPageNumber := range selectedPages {
		sourcePage, ok := sourcePagesMap[oldPageNumber]
		if !ok {
			return nil, nil, fmt.Errorf("%w (%d)", errMissingSourcePage, oldPageNumber)
		}

		if _, ok := forbiddenHashes[sourcePage.FileHash]; ok {
			continue
		}

		if flags.OnlyUniquePages || flags.Only1CopyPages {
			if _, ok := existsPageHashes[sourcePage.FileHash]; ok {
				continue
			}

			existsPageHashes[sourcePage.FileHash] = struct{}{}
		}

		newPageNumberCounter++
		newPageNumber := bookToMerge.PageCount + newPageNumberCounter
		pagesRemap[oldPageNumber] = newPageNumber

		sourcePage.BookID = bookToMerge.ID
		sourcePage.PageNumber = newPageNumber
		sourcePage.CreateAt = time.Now()

		newPages = append(newPages, sourcePage.Page)
	}

	bookToMerge.PageCount += newPageNumberCounter

	return pagesRemap, newPages, nil
}

func (uc *UseCase) rebuildBookSave(
	ctx context.Context,
	isNewBook bool,
	bookToMerge entities.Book,
	newAttributes map[string][]string,
	newPages []entities.Page,
	newLabels []entities.BookLabel,
) (returnErr error) {
	var err error

	if isNewBook {
		err = uc.storage.NewBook(ctx, bookToMerge)
		if err != nil {
			return fmt.Errorf("storage: create book: %w", err)
		}

		defer func() { // Удаляем книгу, если не получилось ее создать полностью
			if returnErr != nil {
				deleteErr := uc.storage.DeleteBook(ctx, bookToMerge.ID)
				if deleteErr != nil {
					uc.logger.ErrorContext(
						ctx, "delete new book after unsuccess rebuild",
						slog.Any("error", deleteErr),
						slog.String("book_id", bookToMerge.ID.String()),
					)
				}
			}
		}()
	} else {
		err = uc.storage.UpdateBook(ctx, bookToMerge)
		if err != nil {
			return fmt.Errorf("storage: create book: %w", err)
		}
	}

	if len(newPages) > 0 {
		err = uc.storage.NewBookPages(ctx, newPages)
		if err != nil {
			return fmt.Errorf("storage: new pages: %w", err)
		}
	}

	if len(newAttributes) > 0 {
		err = uc.storage.UpdateOriginAttributes(ctx, bookToMerge.ID, newAttributes)
		if err != nil {
			return fmt.Errorf("storage: set origin attributes: %w", err)
		}

		err = uc.storage.UpdateAttributes(ctx, bookToMerge.ID, newAttributes)
		if err != nil {
			return fmt.Errorf("storage: set attributes: %w", err)
		}
	}

	if len(newLabels) > 0 {
		err = uc.storage.SetLabels(ctx, newLabels)
		if err != nil {
			return fmt.Errorf("storage: set labels: %w", err)
		}
	}

	return nil
}

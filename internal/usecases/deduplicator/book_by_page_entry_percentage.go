package deduplicator

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) BookByPageEntryPercentage(ctx context.Context, originBookID uuid.UUID) ([]entities.DeduplicateBookResult, error) {
	bookHashes, err := uc.storage.BookPagesWithHash(ctx, originBookID)
	if err != nil {
		return nil, fmt.Errorf("get book hashes storage: %w", err)
	}

	md5Sums := make([]string, len(bookHashes))

	for i, page := range bookHashes {
		md5Sums[i] = page.Md5Sum
	}

	bookIDs, err := uc.storage.BookIDsByMD5(ctx, md5Sums)
	if err != nil {
		return nil, fmt.Errorf("get books by md5 from storage: %w", err)
	}

	bookHandled := make(map[uuid.UUID]struct{}, len(bookIDs))
	bookHandled[originBookID] = struct{}{}

	result := make([]entities.DeduplicateBookResult, 0, len(bookIDs))

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return nil, fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[entities.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	for _, bookID := range bookIDs {
		if _, ok := bookHandled[bookID]; ok {
			continue
		}

		bookHandled[bookID] = struct{}{}

		pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
		if err != nil {
			return nil, fmt.Errorf("get pages (%s) from storage: %w", bookID.String(), err)
		}

		bookShort, err := uc.storage.GetBook(ctx, bookID)
		if err != nil {
			return nil, fmt.Errorf("get book (%s) from storage: %w", bookID.String(), err)
		}

		var previewPage entities.BFFPreviewPage

		for _, page := range pages {
			if page.PageNumber == entities.PageNumberForPreview {
				previewPage = page.ToPreview()
				_, ok := existsDeadHashes[page.FileHash]
				previewPage.HasDeadHash = entities.NewStatusFlag(ok)
			}
		}

		deduplicateResult := entities.DeduplicateBookResult{ // TODO: подумать над оптимизациями
			TargetBook:             bookShort,
			PreviewPage:            previewPage,
			EntryPercentage:        entities.EntryPercentageForPages(bookHashes, pages, nil),
			ReverseEntryPercentage: entities.EntryPercentageForPages(pages, bookHashes, nil),

			EntryPercentageWithoutDeadHashes:        entities.EntryPercentageForPages(bookHashes, pages, existsDeadHashes),
			ReverseEntryPercentageWithoutDeadHashes: entities.EntryPercentageForPages(pages, bookHashes, existsDeadHashes),
		}

		bookHashMap := make(map[entities.FileHash]struct{}, len(bookHashes))

		for _, page := range bookHashes {
			bookHashMap[page.FileHash] = struct{}{}
		}

		for _, page := range pages {
			deduplicateResult.TargetSize.Size += page.Size
			deduplicateResult.TargetSize.Count++

			if _, ok := bookHashMap[page.FileHash]; !ok {
				continue
			}

			deduplicateResult.SharedPages++
			deduplicateResult.SharedSize += page.Size

			if _, ok := existsDeadHashes[page.FileHash]; !ok {
				deduplicateResult.SharedPagesWithoutDeadHashes++
				deduplicateResult.SharedSizeWithoutDeadHashes += page.Size
			}

			delete(bookHashMap, page.FileHash)
		}

		result = append(result, deduplicateResult)
	}

	slices.SortFunc(result, func(a, b entities.DeduplicateBookResult) int {
		aPercent := max(a.EntryPercentage, a.ReverseEntryPercentage)
		bPercent := max(b.EntryPercentage, b.ReverseEntryPercentage)

		if aPercent > bPercent {
			return -1
		}

		if aPercent < bPercent {
			return 1
		}

		if a.TargetSize.Count > b.TargetSize.Count {
			return -1
		}

		if a.TargetSize.Count < b.TargetSize.Count {
			return 1
		}

		if a.EntryPercentage > b.EntryPercentage {
			return -1
		}

		if a.EntryPercentage < b.EntryPercentage {
			return 1
		}

		if a.ReverseEntryPercentage > b.ReverseEntryPercentage {
			return -1
		}

		if a.ReverseEntryPercentage < b.ReverseEntryPercentage {
			return 1
		}

		return 0
	})

	return result, nil
}

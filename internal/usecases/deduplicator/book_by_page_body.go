package deduplicator

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"hgnext/internal/entities"
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

		var previewPage entities.Page

		for _, page := range pages {
			if page.PageNumber == 1 {
				previewPage = page.Page()
			}
		}

		result = append(result, entities.DeduplicateBookResult{
			TargetBook:             bookShort,
			PreviewPage:            previewPage,
			EntryPercentage:        entities.EntryPercentageForPages(bookHashes, pages),
			ReverseEntryPercentage: entities.EntryPercentageForPages(pages, bookHashes),
		})
	}

	slices.SortFunc(result, func(a, b entities.DeduplicateBookResult) int {
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

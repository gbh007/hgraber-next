package deduplicator

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) UniquePages(ctx context.Context, originBookID uuid.UUID) ([]entities.Page, error) {
	originBookPages, err := uc.storage.BookPagesWithHash(ctx, originBookID)
	if err != nil {
		return nil, fmt.Errorf("get book hashes storage: %w", err)
	}

	hashes := make(map[entities.FileHash]struct{}, len(originBookPages))
	md5Sums := make([]string, len(originBookPages))

	for i, page := range originBookPages {
		hashes[page.Hash()] = struct{}{}
		md5Sums[i] = page.Md5Sum
	}

	bookIDs, err := uc.storage.BookIDsByMD5(ctx, md5Sums)
	if err != nil {
		return nil, fmt.Errorf("get books by md5 from storage: %w", err)
	}

	bookHandled := make(map[uuid.UUID]struct{}, len(bookIDs))
	bookHandled[originBookID] = struct{}{}

	for _, bookID := range bookIDs {
		if _, ok := bookHandled[bookID]; ok {
			continue
		}

		bookHandled[bookID] = struct{}{}

		pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
		if err != nil {
			return nil, fmt.Errorf("get pages (%s) from storage: %w", bookID.String(), err)
		}

		for _, page := range pages {
			delete(hashes, page.Hash())
		}
	}

	result := make([]entities.Page, 0, len(hashes))

	for _, page := range originBookPages {
		if _, ok := hashes[page.Hash()]; ok {
			result = append(result, page.Page())
			delete(hashes, page.Hash())
		}
	}

	slices.SortFunc(result, func(a, b entities.Page) int {
		return a.PageNumber - b.PageNumber
	})

	return result, nil
}

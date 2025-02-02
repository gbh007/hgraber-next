package deduplicator

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) UniquePages(ctx context.Context, originBookID uuid.UUID) ([]entities.BFFPreviewPage, error) {
	originBookPages, err := uc.storage.BookPagesWithHash(ctx, originBookID)
	if err != nil {
		return nil, fmt.Errorf("get book hashes storage: %w", err)
	}

	hashes := make(map[entities.FileHash]struct{}, len(originBookPages))
	md5Sums := make([]string, len(originBookPages))

	for i, page := range originBookPages {
		hashes[page.FileHash] = struct{}{}
		md5Sums[i] = page.Md5Sum
	}

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return nil, fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[entities.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	pages, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, md5Sums)
	if err != nil {
		return nil, fmt.Errorf("storage: get pages by md5: %w", err)
	}

	for _, page := range pages {
		if page.BookID == originBookID {
			continue
		}

		delete(hashes, page.FileHash)
	}

	result := make([]entities.BFFPreviewPage, 0, len(hashes))

	for _, page := range originBookPages {
		_, hasDeadHash := existsDeadHashes[page.FileHash]

		if _, ok := hashes[page.FileHash]; ok {
			preview := page.ToPreview()
			preview.HasDeadHash = entities.NewStatusFlag(hasDeadHash)

			result = append(result, preview)

			delete(hashes, page.FileHash)
		}
	}

	slices.SortFunc(result, func(a, b entities.BFFPreviewPage) int {
		return a.PageNumber - b.PageNumber
	})

	return result, nil
}

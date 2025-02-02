package deduplicator

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) DeleteBookDeadHashedPages(ctx context.Context, bookID uuid.UUID) error {
	pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("get book hashes storage: %w", err)
	}

	md5Sums := make([]string, len(pages))
	for i, page := range pages {
		md5Sums[i] = page.Md5Sum
	}

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[entities.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	for _, page := range pages {
		if _, ok := existsDeadHashes[page.FileHash]; !ok {
			continue
		}

		err = uc.storage.MarkPageAsDeleted(ctx, page.BookID, page.PageNumber)
		if err != nil {
			return fmt.Errorf("storage: mark page (%s,%d) as deleted: %w", page.BookID.String(), page.PageNumber, err)
		}
	}

	return nil
}

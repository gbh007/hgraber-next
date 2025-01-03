package deduplicator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) DeleteAllPageByHash(ctx context.Context, bookID uuid.UUID, pageNumber int, setDeadHash bool) error {
	pageHash, err := uc.storage.BookPageWithHash(ctx, bookID, pageNumber)
	if err != nil {
		return fmt.Errorf("storage: get page hash: %w", err)
	}

	pages, err := uc.storage.BookPagesWithHashByHash(ctx, pageHash.Hash())
	if err != nil {
		return fmt.Errorf("storage: get pages by hash: %w", err)
	}

	for _, page := range pages {
		err = uc.storage.MarkPageAsDeleted(ctx, page.BookID, page.PageNumber)
		if err != nil {
			return fmt.Errorf("storage: mark page (%s,%d) as deleted: %w", page.BookID.String(), page.PageNumber, err)
		}
	}

	if setDeadHash {
		err = uc.storage.SetDeadHash(ctx, entities.DeadHash{
			FileHash:  pageHash.Hash(),
			CreatedAt: time.Now().UTC(),
		})
		if err != nil {
			return fmt.Errorf("storage: set dead hash: %w", err)
		}
	}

	return nil
}

package deduplicator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) CreateDeadHashByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) error {
	pageHash, err := uc.storage.BookPageWithHash(ctx, bookID, pageNumber)
	if err != nil {
		return fmt.Errorf("storage: get page hash: %w", err)
	}

	err = uc.storage.SetDeadHash(ctx, entities.DeadHash{
		FileHash:  pageHash.Hash(),
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("storage: set dead hash: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteDeadHashByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) error {
	pageHash, err := uc.storage.BookPageWithHash(ctx, bookID, pageNumber)
	if err != nil {
		return fmt.Errorf("storage: get page hash: %w", err)
	}

	err = uc.storage.DeleteDeadHash(ctx, entities.DeadHash{FileHash: pageHash.Hash()})
	if err != nil {
		return fmt.Errorf("storage: delete dead hash: %w", err)
	}

	return nil
}

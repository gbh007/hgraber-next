package deduplicator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) MarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error {
	pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get pages: %w", err)
	}

	for _, page := range pages {
		err = uc.storage.SetDeadHash(ctx, entities.DeadHash{
			FileHash:  page.FileHash,
			CreatedAt: time.Now().UTC(),
		})
		if err != nil {
			return fmt.Errorf("storage: set dead hash (%d): %w", page.PageNumber, err)
		}
	}

	return nil
}

func (uc *UseCase) UnMarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error {
	pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get pages: %w", err)
	}

	for _, page := range pages {
		err = uc.storage.DeleteDeadHash(ctx, entities.DeadHash{
			FileHash: page.FileHash,
		})
		if err != nil {
			return fmt.Errorf("storage: delete dead hash (%d): %w", page.PageNumber, err)
		}
	}

	return nil
}

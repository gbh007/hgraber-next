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

func (uc *UseCase) RemoveBookPagesWithDeadHash(ctx context.Context, bookID uuid.UUID, deleteEmptyBook bool) error {
	masterPages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get pages: %w", err)
	}

	bookIDMap := make(map[uuid.UUID]struct{})

	for _, masterPage := range masterPages {
		err = uc.storage.SetDeadHash(ctx, entities.DeadHash{
			FileHash:  masterPage.FileHash,
			CreatedAt: time.Now().UTC(),
		})
		if err != nil {
			return fmt.Errorf("storage: set dead hash (%d): %w", masterPage.PageNumber, err)
		}

		pages, err := uc.storage.BookPagesWithHashByHash(ctx, masterPage.FileHash)
		if err != nil {
			return fmt.Errorf("storage: get pages by hash (%d): %w", masterPage.PageNumber, err)
		}

		for _, page := range pages {
			bookIDMap[page.BookID] = struct{}{}

			err = uc.storage.MarkPageAsDeleted(ctx, page.BookID, page.PageNumber)
			if err != nil {
				return fmt.Errorf("storage: mark page (%s,%d) as deleted: %w", page.BookID.String(), page.PageNumber, err)
			}
		}
	}

	if !deleteEmptyBook {
		return nil
	}

	for bookID := range bookIDMap {
		count, err := uc.storage.BookPagesCount(ctx, bookID)
		if err != nil {
			return fmt.Errorf("storage: get book page count (%s): %w", bookID.String(), err)
		}

		if count != 0 {
			continue
		}

		err = uc.storage.MarkBookAsDeleted(ctx, bookID)
		if err != nil {
			return fmt.Errorf("storage: mark book as deleted (%s): %w", bookID.String(), err)
		}
	}

	return nil
}

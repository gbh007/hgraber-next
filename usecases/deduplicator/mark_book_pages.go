package deduplicator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) MarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error {
	pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get pages: %w", err)
	}

	deadHashes := make([]core.DeadHash, 0, len(pages))

	for _, page := range pages {
		deadHashes = append(deadHashes, core.DeadHash{
			FileHash:  page.FileHash,
			CreatedAt: time.Now().UTC(),
		})
	}

	err = uc.storage.SetDeadHashes(ctx, deadHashes)
	if err != nil {
		return fmt.Errorf("storage: set dead hashes: %w", err)
	}

	return nil
}

func (uc *UseCase) UnMarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error {
	pages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get pages: %w", err)
	}

	deadHashes := make([]core.DeadHash, 0, len(pages))

	for _, page := range pages {
		deadHashes = append(deadHashes, core.DeadHash{
			FileHash:  page.FileHash,
			CreatedAt: time.Now().UTC(),
		})
	}

	err = uc.storage.DeleteDeadHashes(ctx, deadHashes)
	if err != nil {
		return fmt.Errorf("storage: delete dead hashes: %w", err)
	}

	return nil
}

func (uc *UseCase) RemoveBookPagesWithDeadHash(ctx context.Context, bookID uuid.UUID, deleteEmptyBook bool) error {
	masterPages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get pages: %w", err)
	}

	md5Sums := make([]string, len(masterPages))
	masterPageHashes := make(map[core.FileHash]struct{}, len(masterPages))
	deadHashes := make([]core.DeadHash, 0, len(masterPages))

	for _, page := range masterPages {
		masterPageHashes[page.FileHash] = struct{}{}

		md5Sums = append(md5Sums, page.Md5Sum)
		deadHashes = append(deadHashes, core.DeadHash{
			FileHash:  page.FileHash,
			CreatedAt: time.Now().UTC(),
		})
	}

	err = uc.storage.SetDeadHashes(ctx, deadHashes)
	if err != nil {
		return fmt.Errorf("storage: set dead hashes: %w", err)
	}

	pages, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, md5Sums)
	if err != nil {
		return fmt.Errorf("storage: get pages by md5: %w", err)
	}

	bookIDMap := make(map[uuid.UUID]struct{})

	for _, page := range pages {
		// Отсекаем не совпадения из-за неполного ограничения в БД
		if _, ok := masterPageHashes[page.FileHash]; !ok {
			continue
		}

		bookIDMap[page.BookID] = struct{}{}

		err = uc.storage.MarkPageAsDeleted(ctx, page.BookID, page.PageNumber)
		if err != nil {
			return fmt.Errorf("storage: mark page (%s,%d) as deleted: %w", page.BookID.String(), page.PageNumber, err)
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

package rebuilder

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) rebuildBookCleanSource(
	ctx context.Context,
	flags entities.RebuildBookRequestFlags,
	sourceBookID uuid.UUID,
	unusedSourceHashes map[entities.FileHash]struct{},
	usedSourcePageNumbers []int,
) error {
	if flags.MarkUnusedPagesAsDeadHash && len(unusedSourceHashes) > 0 {
		deadHashes := make([]entities.DeadHash, 0, len(unusedSourceHashes))

		for hash := range unusedSourceHashes {
			deadHashes = append(deadHashes, entities.DeadHash{
				FileHash:  hash,
				CreatedAt: time.Now().UTC(),
			})
		}

		err := uc.storage.SetDeadHashes(ctx, deadHashes)
		if err != nil {
			return fmt.Errorf("storage: set dead hashes: %w", err)
		}
	}

	if flags.MarkUnusedPagesAsDeleted && len(unusedSourceHashes) > 0 {
		md5Sums := make([]string, len(unusedSourceHashes))

		for hash := range unusedSourceHashes {
			md5Sums = append(md5Sums, hash.Md5Sum)
		}

		pages, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, md5Sums)
		if err != nil {
			return fmt.Errorf("storage: get pages by md5: %w", err)
		}

		bookIDs := make(map[uuid.UUID]struct{})

		for _, page := range pages {
			// Отсекаем не совпадения из-за неполного ограничения в БД
			if _, ok := unusedSourceHashes[page.FileHash]; !ok {
				continue
			}

			bookIDs[page.BookID] = struct{}{}

			err = uc.storage.MarkPageAsDeleted(ctx, page.BookID, page.PageNumber)
			if err != nil {
				return fmt.Errorf("storage: mark page (%s,%d) as deleted: %w", page.BookID.String(), page.PageNumber, err)
			}
		}

		if flags.MarkEmptyBookAsDeletedAfterRemovePages {
			for bookID := range bookIDs {
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
		}
	}

	if flags.ExtractMode && len(usedSourcePageNumbers) > 0 {
		for _, pageNUmber := range usedSourcePageNumbers {
			err := uc.storage.MarkPageAsDeleted(ctx, sourceBookID, pageNUmber)
			if err != nil {
				return fmt.Errorf("storage: mark page (%s,%d) as deleted: %w", sourceBookID.String(), pageNUmber, err)
			}
		}

		if flags.MarkEmptyBookAsDeletedAfterRemovePages {
			count, err := uc.storage.BookPagesCount(ctx, sourceBookID)
			if err != nil {
				return fmt.Errorf("storage: get book page count (%s): %w", sourceBookID.String(), err)
			}

			if count == 0 {
				err = uc.storage.MarkBookAsDeleted(ctx, sourceBookID)
				if err != nil {
					return fmt.Errorf("storage: mark book as deleted (%s): %w", sourceBookID.String(), err)
				}
			}
		}
	}

	return nil
}

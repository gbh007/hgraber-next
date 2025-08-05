package cleanupusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func (uc *UseCase) CleanAfterRebuild(_ context.Context) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("CleanupAfterRebuild")

		taskResult.StartStage("search hashes")

		deletedPagesHashes, err := uc.storage.DeletedPagesHashes(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(deletedPagesHashes)))
		taskResult.SetProgress(int64(len(deletedPagesHashes)))
		taskResult.EndStage()

		taskResult.StartStage("search dead hashes")
		taskResult.SetTotal(int64(len(deletedPagesHashes)))

		deadHashes := make([]core.DeadHash, 0, len(deletedPagesHashes))
		deadHashesToRemoveFromDeletedPages := make([]core.FileHash, 0, len(deletedPagesHashes))
		existsHashesToRemoveFromDeletedPages := make([]core.FileHash, 0, len(deletedPagesHashes))

		for _, hash := range deletedPagesHashes {
			taskResult.IncProgress()

			pageCount, err := uc.storage.BookPagesCountByHash(ctx, hash)
			if err != nil {
				taskResult.SetError(fmt.Errorf("get page count by hash: %w", err))

				return
			}

			if pageCount > 0 { // Есть активные страницы
				existsHashesToRemoveFromDeletedPages = append(existsHashesToRemoveFromDeletedPages, hash)

				continue
			}

			deadHashes = append(deadHashes, core.DeadHash{
				FileHash:  hash,
				CreatedAt: time.Now().UTC(),
			})

			deadHashesToRemoveFromDeletedPages = append(deadHashesToRemoveFromDeletedPages, hash)
		}

		taskResult.EndStage()

		taskResult.StartStage("set dead hashes")
		taskResult.SetTotal(int64(len(deadHashes)))

		if len(deadHashes) > 0 {
			err = uc.storage.SetDeadHashes(ctx, deadHashes)
			if err != nil {
				taskResult.SetError(fmt.Errorf("set dead hashes: %w", err))

				return
			}
		}

		taskResult.SetProgress(int64(len(deadHashes)))
		taskResult.EndStage()

		taskResult.StartStage("remove deleted page by new dead hashes")
		taskResult.SetTotal(int64(len(deadHashesToRemoveFromDeletedPages)))

		if len(deadHashesToRemoveFromDeletedPages) > 0 {
			err = uc.storage.RemoveDeletedPagesByHashes(ctx, deadHashesToRemoveFromDeletedPages)
			if err != nil {
				taskResult.SetError(fmt.Errorf("remove deleted page by new dead hashes: %w", err))

				return
			}
		}

		taskResult.SetProgress(int64(len(deadHashesToRemoveFromDeletedPages)))
		taskResult.EndStage()

		taskResult.StartStage("remove deleted page by exists hashes")
		taskResult.SetTotal(int64(len(existsHashesToRemoveFromDeletedPages)))

		if len(existsHashesToRemoveFromDeletedPages) > 0 {
			err = uc.storage.RemoveDeletedPagesByHashes(ctx, existsHashesToRemoveFromDeletedPages)
			if err != nil {
				taskResult.SetError(fmt.Errorf("remove deleted page by exists hashes: %w", err))

				return
			}
		}

		taskResult.SetProgress(int64(len(existsHashesToRemoveFromDeletedPages)))
		taskResult.EndStage()

		detachedFileCount, detachedFileSize, err := uc.removeDetachedFiles(ctx, taskResult)
		if err != nil {
			return // В результат выполнения данные уже проставлены
		}

		deletedRebuildIDs, err := uc.cleanDeletedRebuilds(ctx, taskResult)
		if err != nil {
			return // В результат выполнения данные уже проставлены
		}

		taskResult.SetResult(fmt.Sprintf(
			"Новых мертвых хешей %d\n"+
				"Удаленные файлы = количество: %d байт: %d размер: %s\n"+
				"Удалено удаленных ребилдов: %d",
			len(deadHashes),
			detachedFileCount,
			detachedFileSize,
			core.PrettySize(detachedFileSize),
			len(deletedRebuildIDs),
		))
	}), nil
}

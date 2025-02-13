package deduplicatorusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func (uc *UseCase) FillDeadHashes(_ context.Context, withRemoveDeletedPages bool) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		if withRemoveDeletedPages {
			taskResult.SetName("FillDeadHashesWithRemoveDeletedPages")
		} else {
			taskResult.SetName("FillDeadHashes")
		}

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

		for _, hash := range deletedPagesHashes {
			taskResult.IncProgress()

			pageCount, err := uc.storage.BookPagesCountByHash(ctx, hash)
			if err != nil {
				taskResult.SetError(fmt.Errorf("get page count by hash: %w", err))

				return
			}

			if pageCount > 0 { // Есть активные страницы
				continue
			}

			deadHashes = append(deadHashes, core.DeadHash{
				FileHash:  hash,
				CreatedAt: time.Now().UTC(),
			})

			if withRemoveDeletedPages {
				deadHashesToRemoveFromDeletedPages = append(deadHashesToRemoveFromDeletedPages, hash)
			}
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

		if withRemoveDeletedPages {
			taskResult.StartStage("remove deleted page by new dead hashes")
			taskResult.SetTotal(int64(len(deadHashesToRemoveFromDeletedPages)))

			if len(deadHashesToRemoveFromDeletedPages) > 0 {
				err = uc.storage.RemoveDeletedPagesByHashes(ctx, deadHashesToRemoveFromDeletedPages)
				if err != nil {
					taskResult.SetError(fmt.Errorf("remove deleted pages by hashes: %w", err))

					return
				}
			}

			taskResult.SetProgress(int64(len(deadHashesToRemoveFromDeletedPages)))
			taskResult.EndStage()
		}

		taskResult.SetResult(fmt.Sprintf("Обработано %d", len(deadHashes)))
	}), nil
}

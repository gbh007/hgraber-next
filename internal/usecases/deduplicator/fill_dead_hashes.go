package deduplicator

import (
	"context"
	"fmt"
	"time"

	"hgnext/internal/entities"
)

func (uc *UseCase) FillDeadHashes(_ context.Context) (entities.RunnableTask, error) {
	return entities.RunnableTaskFunction(func(ctx context.Context, taskResult entities.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("FillDeadHashes")

		taskResult.StartStage("search hashes")

		deletedPagesHashes, err := uc.storage.DeletedPagesHashes(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(deletedPagesHashes)))
		taskResult.SetProgress(int64(len(deletedPagesHashes)))
		taskResult.EndStage()

		taskResult.StartStage("fill dead hashes")
		taskResult.SetTotal(int64(len(deletedPagesHashes)))

		count := 0

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

			err = uc.storage.SetDeadHash(ctx, entities.DeadHash{
				FileHash:  hash,
				CreatedAt: time.Now().UTC(),
			})
			if err != nil {
				taskResult.SetError(fmt.Errorf("set dead hash: %w", err))

				return
			}

			count++
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf("Обработано %d", count))
	}), nil
}

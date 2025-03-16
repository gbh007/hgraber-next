package attributeusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func (uc *UseCase) RemapBooks(_ context.Context) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("RemapBooks")

		remaps, err := uc.storage.AttributeRemaps(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		remaper := core.NewAttributeRemaper(remaps, uc.remapToLower)

		taskResult.StartStage("search books")

		ids, err := uc.storage.BookIDs(ctx, core.BookFilter{
			ShowDeleted: core.BookFilterShowTypeExcept,
		})
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(ids)))
		taskResult.SetProgress(int64(len(ids)))
		taskResult.EndStage()

		taskResult.StartStage("update books")
		taskResult.SetTotal(int64(len(ids)))

		for _, bookID := range ids {
			taskResult.IncProgress()

			attributes, err := uc.storage.BookOriginAttributes(ctx, bookID)
			if err != nil {
				taskResult.SetError(err)

				return
			}

			attributes = remaper.Remap(attributes)

			if len(attributes) > 0 {
				err = uc.storage.UpdateAttributes(ctx, bookID, attributes)
				if err != nil {
					taskResult.SetError(err)

					return
				}
			} else {
				err = uc.storage.DeleteBookAttributes(ctx, bookID)
				if err != nil {
					taskResult.SetError(err)

					return
				}
			}
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf("Обновлено %d", len(ids)))
	}), nil
}

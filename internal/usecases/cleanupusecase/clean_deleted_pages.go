package cleanupusecase

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/systemmodel"
)

func (uc *UseCase) CleanDeletedPages(_ context.Context) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("CleanDeletedPages")

		taskResult.StartStage("truncate deleted pages")

		err := uc.storage.TruncateDeletedPages(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.EndStage()
	}), nil
}

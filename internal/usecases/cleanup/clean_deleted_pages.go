package cleanup

import (
	"context"

	"hgnext/internal/entities"
)

func (uc *UseCase) CleanDeletedPages(_ context.Context) (entities.RunnableTask, error) {
	return entities.RunnableTaskFunction(func(ctx context.Context, taskResult entities.TaskResultWriter) {
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

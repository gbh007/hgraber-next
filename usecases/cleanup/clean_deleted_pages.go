package cleanup

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) CleanDeletedPages(_ context.Context) (core.RunnableTask, error) {
	return core.RunnableTaskFunction(func(ctx context.Context, taskResult core.TaskResultWriter) {
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

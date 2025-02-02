package cleanup

import (
	"bytes"
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) CleanDeletedRebuilds(_ context.Context) (entities.RunnableTask, error) {
	return entities.RunnableTaskFunction(func(ctx context.Context, taskResult entities.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("CleanDeletedRebuilds")

		taskResult.StartStage("search deleted rebuilds")

		ids, err := uc.storage.BookIDsWithDeletedRebuilds(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		buff := &bytes.Buffer{}

		for _, id := range ids {
			buff.WriteString(id.String() + "\n")
		}

		taskResult.SetResult(buff.String())
		taskResult.EndStage()

		taskResult.StartStage("remove deleted rebuilds")

		if len(ids) > 0 {
			err = uc.storage.DeleteBooks(ctx, ids)
			if err != nil {
				taskResult.SetError(err)

				return
			}
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf("Удалено %d", len(ids)))
	}), nil
}

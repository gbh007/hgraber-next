package cleanup

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
)

func (uc *UseCase) RemoveDetachedFiles(_ context.Context) (entities.RunnableTask, error) {
	return entities.RunnableTaskFunction(func(ctx context.Context, taskResult entities.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("RemoveDetachedFiles")

		ctx, span := uc.tracer.Start(ctx, "RemoveDetachedFiles")
		defer span.End()

		taskResult.StartStage("search detached files")
		span.AddEvent("search detached files", trace.WithTimestamp(time.Now()))

		files, err := uc.storage.DetachedFiles(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(files)))
		taskResult.SetProgress(int64(len(files)))
		taskResult.EndStage()

		taskResult.StartStage("remove detached files")
		taskResult.SetTotal(int64(len(files)))
		span.AddEvent("remove detached files", trace.WithTimestamp(time.Now()))

		var (
			count int
			size  int64
		)

		for _, file := range files {
			taskResult.IncProgress()

			if file.Size == 0 {
				uc.logger.WarnContext(
					ctx, "empty file size",
					slog.String("id", file.ID.String()),
				)

				continue
			}

			err = uc.storage.DeleteFile(ctx, file.ID)
			if err != nil {
				taskResult.SetError(fmt.Errorf("delete file (%s) from storage: %w", file.ID.String(), err))

				return
			}

			err = uc.fileStorage.Delete(ctx, file.ID)
			if err != nil {
				taskResult.SetError(fmt.Errorf("delete file (%s) from file-storage: %w", file.ID.String(), err))

				return
			}

			count++

			size += file.Size
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf(
			"count: %d size: %d human size: %s",
			count, size, entities.PrettySize(size),
		))
	}), nil
}

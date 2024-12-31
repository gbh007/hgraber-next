package deduplicator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
)

type fileKey struct {
	md5    string
	sha256 string
	size   int64
}

func (uc *UseCase) DeduplicateFiles(_ context.Context) (entities.RunnableTask, error) {
	return entities.RunnableTaskFunction(func(ctx context.Context, taskResult entities.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("DeduplicateFiles")

		ctx, span := uc.tracer.Start(ctx, "DeduplicateFiles")
		defer span.End()

		taskResult.StartStage("get duplicates from storage")
		span.AddEvent("get duplicates from storage", trace.WithTimestamp(time.Now()))

		files, err := uc.storage.DuplicatedFiles(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(files)))
		taskResult.SetProgress(int64(len(files)))
		taskResult.EndStage()

		taskResult.StartStage("transform data")
		taskResult.SetTotal(int64(len(files)))
		span.AddEvent("transform data", trace.WithTimestamp(time.Now()))

		fileMap := make(map[fileKey][]uuid.UUID)

		for _, file := range files {
			taskResult.IncProgress()

			k := fileKey{
				md5:    file.Md5Sum,
				sha256: file.Sha256Sum,
				size:   file.Size,
			}

			fileMap[k] = append(fileMap[k], file.ID)
		}

		taskResult.EndStage()

		taskResult.StartStage("handle duplicates")
		taskResult.SetTotal(int64(len(fileMap)))
		span.AddEvent("handle duplicates", trace.WithTimestamp(time.Now()))

		var (
			count int
			size  int64
		)

		for k, ids := range fileMap {
			taskResult.IncProgress()

			if k.size == 0 {
				uc.logger.WarnContext(
					ctx, "empty file size",
					slog.Any("ids", ids),
				)

				continue
			}

			if len(ids) < 2 {
				uc.logger.WarnContext(
					ctx, "invalid deduplicate ids len",
					slog.Any("ids", ids),
				)

				continue
			}

			newID := ids[0]
			ids = ids[1:]

			for _, id := range ids {
				err = uc.storage.ReplaceFile(ctx, id, newID)
				if err != nil {
					taskResult.SetError(fmt.Errorf("replace id in storage: %w", err))

					return
				}

				uc.logger.InfoContext(
					ctx, "replaced file",
					slog.String("old", id.String()),
					slog.String("new", newID.String()),
				)
			}

			size += k.size * int64(len(ids))
			count += len(ids)
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf(
			"count: %d size: %d human size: %s",
			count, size, entities.PrettySize(size),
		))
	}), nil
}

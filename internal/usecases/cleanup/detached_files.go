package cleanup

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/trace"
)

func (uc *UseCase) RemoveDetachedFiles(ctx context.Context) (count int, size int64, err error) {
	ctx, span := uc.tracer.Start(ctx, "RemoveDetachedFiles")
	defer span.End()

	span.AddEvent("search detached files", trace.WithTimestamp(time.Now()))

	files, err := uc.storage.DetachedFiles(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("get detached files from storage: %w", err)
	}

	span.AddEvent("remove detached files", trace.WithTimestamp(time.Now()))

	for _, file := range files {
		if file.Size == 0 {
			uc.logger.Logger(ctx).WarnContext(
				ctx, "empty file size",
				slog.String("id", file.ID.String()),
			)

			continue
		}

		err = uc.storage.DeleteFile(ctx, file.ID)
		if err != nil {
			return 0, 0, fmt.Errorf("delete file (%s) from storage: %w", file.ID.String(), err)
		}

		err = uc.fileStorage.Delete(ctx, file.ID)
		if err != nil {
			return 0, 0, fmt.Errorf("delete file (%s) from file-storage: %w", file.ID.String(), err)
		}

		count++

		size += file.Size
	}

	return count, size, nil
}

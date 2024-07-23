package deduplicator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type fileKey struct {
	md5    string
	sha256 string
	size   int64
}

func (uc *UseCase) DeduplicateFiles(ctx context.Context) (count int, size int64, err error) {
	ctx, span := uc.tracer.Start(ctx, "DeduplicateFiles")
	defer span.End()

	span.AddEvent("get duplicates from storage", trace.WithTimestamp(time.Now()))

	files, err := uc.storage.DuplicatedFiles(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("get duplicates from storage: %w", err)
	}

	span.AddEvent("transform data", trace.WithTimestamp(time.Now()))

	fileMap := make(map[fileKey][]uuid.UUID)

	for _, file := range files {
		k := fileKey{
			md5:    file.Md5Sum,
			sha256: file.Sha256Sum,
			size:   file.Size,
		}

		fileMap[k] = append(fileMap[k], file.ID)
	}

	span.AddEvent("handle duplicates", trace.WithTimestamp(time.Now()))

	for k, ids := range fileMap {
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
				return 0, 0, fmt.Errorf("replace id in storage: %w", err)
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

	return count, size, nil
}

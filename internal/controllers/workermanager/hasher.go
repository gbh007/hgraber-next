package workermanager

import (
	"context"
	"log/slog"
	"time"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
)

type hasherUnitUseCases interface {
	UnHashedFiles(ctx context.Context) ([]entities.File, error)
	HandleFileHash(ctx context.Context, f entities.File) error
}

func NewHasher(useCases hasherUnitUseCases, logger *slog.Logger) *worker.Worker[entities.File] {
	return worker.New[entities.File](
		"file_hash",
		1000,
		time.Second*15,
		logger,
		func(ctx context.Context, file entities.File) {
			err := useCases.HandleFileHash(ctx, file)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail hash file",
					slog.String("file_id", file.ID.String()),
					slog.Any("error", err),
				)
			}
		},
		func(ctx context.Context) []entities.File {
			files, err := useCases.UnHashedFiles(ctx)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail get files for hashing",
					slog.Any("error", err),
				)
			}

			return files
		},
		10,
	)
}

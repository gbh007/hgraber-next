package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type hasherUnitUseCases interface {
	UnHashedFiles(ctx context.Context) ([]core.File, error)
	HandleFileHash(ctx context.Context, f core.File) error
}

func NewHasher(
	useCases hasherUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[core.File] {
	return worker.New[core.File](
		"file_hash",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, file core.File) error {
			err := useCases.HandleFileHash(ctx, file)
			if err != nil {
				return pkg.WrapError(
					err, "fail hash file",
					pkg.ErrorArgument("file_id", file.ID),
				)
			}

			return nil
		},
		useCases.UnHashedFiles,
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/pkg"
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
		cfg,
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
		tracer,
		metricProvider,
	)
}

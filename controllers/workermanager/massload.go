package workermanager

import (
	"context"
	"log/slog"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/pkg"
	"go.opentelemetry.io/otel/trace"
)

type massLoadSizeUnitUseCases interface {
	MassloadForUpdate(ctx context.Context) ([]massloadmodel.Massload, error)
	UpdateSize(ctx context.Context, ml massloadmodel.Massload) error
}

func NewMassLoadSize(
	useCases massLoadSizeUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Massload] {
	return worker.New[massloadmodel.Massload](
		"massload_sizer",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, ml massloadmodel.Massload) error {
			err := useCases.UpdateSize(ctx, ml)
			if err != nil {
				return pkg.WrapError(
					err, "fail update massload size info",
					pkg.ErrorArgument("massload_id", ml.ID),
				)
			}

			return nil
		},
		useCases.MassloadForUpdate,
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

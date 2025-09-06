package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

type massloadCalculationUnitUseCases interface {
	MassloadForUpdateCalculation(ctx context.Context) ([]massloadmodel.Massload, error)
	UpdateCalculation(ctx context.Context, ml massloadmodel.Massload) error
}

func NewMassloadCalculation(
	useCases massloadCalculationUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Massload] {
	return worker.New[massloadmodel.Massload](
		cfg,
		logger,
		func(ctx context.Context, ml massloadmodel.Massload) error {
			err := useCases.UpdateCalculation(ctx, ml)
			if err != nil {
				return pkg.WrapError(
					err, "fail update massload calculation info",
					pkg.ErrorArgument("massload_id", ml.ID),
				)
			}

			return nil
		},
		useCases.MassloadForUpdateCalculation,
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

type massloadSizeUnitUseCases interface {
	MassloadForUpdate(ctx context.Context) ([]massloadmodel.Massload, error)
	UpdateSize(ctx context.Context, ml massloadmodel.Massload) error
}

type massloadAttributeSizeUnitUseCases interface {
	MassloadAttributesForUpdate(ctx context.Context) ([]massloadmodel.Attribute, error)
	UpdateAttributesSize(ctx context.Context, attr massloadmodel.Attribute) error
}

func NewMassloadSize(
	useCases massloadSizeUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Massload] {
	return worker.New[massloadmodel.Massload](
		cfg,
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
		tracer,
		metricProvider,
	)
}

func NewMassloadAttributeSize(
	useCases massloadAttributeSizeUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Attribute] {
	return worker.New[massloadmodel.Attribute](
		cfg,
		logger,
		func(ctx context.Context, attr massloadmodel.Attribute) error {
			err := useCases.UpdateAttributesSize(ctx, attr)
			if err != nil {
				return pkg.WrapError(
					err, "fail update massload attribute size info",
					pkg.ErrorArgument("code", attr.Code),
					pkg.ErrorArgument("value", attr.Value),
				)
			}

			return nil
		},
		useCases.MassloadAttributesForUpdate,
		tracer,
		metricProvider,
	)
}

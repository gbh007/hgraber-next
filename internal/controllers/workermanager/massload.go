package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/massloadmodel"
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
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Massload] {
	return worker.New[massloadmodel.Massload](
		cfg,
		logger,
		func(ctx context.Context, ml massloadmodel.Massload) error {
			err := useCases.UpdateSize(ctx, ml)
			if err != nil {
				return fmt.Errorf("fail update massload size info (massload_id: %d): %w", ml.ID, err)
			}

			return nil
		},
		useCases.MassloadForUpdate,
		metricProvider,
	)
}

func NewMassloadAttributeSize(
	useCases massloadAttributeSizeUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Attribute] {
	return worker.New[massloadmodel.Attribute](
		cfg,
		logger,
		func(ctx context.Context, attr massloadmodel.Attribute) error {
			err := useCases.UpdateAttributesSize(ctx, attr)
			if err != nil {
				return fmt.Errorf(
					"fail update massload attribute size info (code: %s, value: %s): %w",
					attr.Code, attr.Value, err,
				)
			}

			return nil
		},
		useCases.MassloadAttributesForUpdate,
		metricProvider,
	)
}

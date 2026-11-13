package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/massloadmodel"
)

type massloadCalculationUnitUseCases interface {
	MassloadForUpdateCalculation(ctx context.Context) ([]massloadmodel.Massload, error)
	UpdateCalculation(ctx context.Context, ml massloadmodel.Massload) error
}

func NewMassloadCalculation(
	useCases massloadCalculationUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[massloadmodel.Massload] {
	return worker.New[massloadmodel.Massload](
		cfg,
		logger,
		func(ctx context.Context, ml massloadmodel.Massload) error {
			err := useCases.UpdateCalculation(ctx, ml)
			if err != nil {
				return fmt.Errorf("fail update massload calculation info (massload_id: %d): %w", ml.ID, err)
			}

			return nil
		},
		useCases.MassloadForUpdateCalculation,
		metricProvider,
	)
}

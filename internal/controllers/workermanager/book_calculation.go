package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
)

type bookCalculationUnitUseCases interface {
	GetBookIDsForCalculation(ctx context.Context) ([]uuid.UUID, error)
	CalculateBook(ctx context.Context, bookID uuid.UUID) error
}

func NewBookCalculation(
	useCases bookCalculationUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[uuid.UUID] {
	return worker.New[uuid.UUID](
		cfg,
		logger,
		func(ctx context.Context, id uuid.UUID) error {
			err := useCases.CalculateBook(ctx, id)
			if err != nil {
				return fmt.Errorf("fail update book calculation (book_id: %s): %w", id.String(), err)
			}

			return nil
		},
		useCases.GetBookIDsForCalculation,
		metricProvider,
	)
}

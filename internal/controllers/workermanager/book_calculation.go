package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/pkg"
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
				return pkg.WrapError(
					err, "fail update book calculation",
					pkg.ErrorArgument("book_id", id.String()),
				)
			}

			return nil
		},
		useCases.GetBookIDsForCalculation,
		metricProvider,
	)
}

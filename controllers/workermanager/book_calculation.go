package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/pkg"
)

type bookCalculationUnitUseCases interface {
	GetBookIDsForCalculation(ctx context.Context) ([]uuid.UUID, error)
	CalculateBook(ctx context.Context, bookID uuid.UUID) error
}

func NewBookCalculation(
	useCases bookCalculationUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
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
		tracer,
		metricProvider,
	)
}

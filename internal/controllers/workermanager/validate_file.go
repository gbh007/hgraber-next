package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/controllers/internal/worker"
)

type validateFileUnitUseCases interface {
	ValidateFile(ctx context.Context, fileID uuid.UUID) error
	FileIDsToValidate() []uuid.UUID
}

func NewFileValidator(
	useCases validateFileUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[uuid.UUID] {
	return worker.New[uuid.UUID](
		"validate_file",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, fileID uuid.UUID) {
			err := useCases.ValidateFile(ctx, fileID)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail validate file",
					slog.String("id", fileID.String()),
					slog.Any("error", err),
				)
			}
		},
		func(_ context.Context) []uuid.UUID {
			return useCases.FileIDsToValidate()
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

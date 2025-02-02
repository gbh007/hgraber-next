package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/internal/pkg"
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
		func(ctx context.Context, fileID uuid.UUID) error {
			err := useCases.ValidateFile(ctx, fileID)
			if err != nil {
				return pkg.WrapError(
					err, "fail validate file",
					pkg.ErrorArgument("file_id", fileID),
				)
			}

			return nil
		},
		func(_ context.Context) ([]uuid.UUID, error) {
			return useCases.FileIDsToValidate(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
)

type validateFileUnitUseCases interface {
	ValidateFile(ctx context.Context, fileID uuid.UUID) error
	FileIDsToValidate() []uuid.UUID
}

func NewFileValidator(
	useCases validateFileUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[uuid.UUID] {
	return worker.New[uuid.UUID](
		cfg,
		logger,
		func(ctx context.Context, fileID uuid.UUID) error {
			err := useCases.ValidateFile(ctx, fileID)
			if err != nil {
				return fmt.Errorf("fail validate file (file_id: %s): %w", fileID.String(), err)
			}

			return nil
		},
		func(_ context.Context) ([]uuid.UUID, error) {
			return useCases.FileIDsToValidate(), nil
		},
		metricProvider,
	)
}

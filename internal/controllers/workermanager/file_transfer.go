package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/fsmodel"
)

type fileTransferUnitUseCases interface {
	TransferFile(ctx context.Context, transfer fsmodel.FileTransfer) error
	FileTransferList() []fsmodel.FileTransfer
}

func NewFileTransfer(
	useCases fileTransferUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[fsmodel.FileTransfer] {
	return worker.New[fsmodel.FileTransfer](
		cfg,
		logger,
		func(ctx context.Context, transfer fsmodel.FileTransfer) error {
			err := useCases.TransferFile(ctx, transfer)
			if err != nil {
				return fmt.Errorf(
					"fail transfer file (file_id: %s, fs_id: %s): %w",
					transfer.FileID.String(), transfer.FSID.String(), err,
				)
			}

			return nil
		},
		func(_ context.Context) ([]fsmodel.FileTransfer, error) {
			return useCases.FileTransferList(), nil
		},
		metricProvider,
	)
}

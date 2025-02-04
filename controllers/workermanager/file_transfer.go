package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/pkg"
)

type fileTransferUnitUseCases interface {
	TransferFile(ctx context.Context, transfer entities.FileTransfer) error
	FileTransferList() []entities.FileTransfer
}

func NewFileTransfer(
	useCases fileTransferUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[entities.FileTransfer] {
	return worker.New[entities.FileTransfer](
		"transfer_file",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, transfer entities.FileTransfer) error {
			err := useCases.TransferFile(ctx, transfer)
			if err != nil {
				return pkg.WrapError(
					err, "fail transfer file",
					pkg.ErrorArgument("file_id", transfer.FileID),
					pkg.ErrorArgument("fs_id", transfer.FSID),
				)
			}

			return nil
		},
		func(_ context.Context) ([]entities.FileTransfer, error) {
			return useCases.FileTransferList(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
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
		func(ctx context.Context, transfer entities.FileTransfer) {
			err := useCases.TransferFile(ctx, transfer)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail transfer file",
					slog.String("file_id", transfer.FileID.String()),
					slog.String("fs_id", transfer.FSID.String()),
					slog.Any("error", err),
				)
			}
		},
		func(_ context.Context) []entities.FileTransfer {
			return useCases.FileTransferList()
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

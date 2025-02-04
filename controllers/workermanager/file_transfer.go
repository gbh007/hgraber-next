package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type fileTransferUnitUseCases interface {
	TransferFile(ctx context.Context, transfer core.FileTransfer) error
	FileTransferList() []core.FileTransfer
}

func NewFileTransfer(
	useCases fileTransferUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[core.FileTransfer] {
	return worker.New[core.FileTransfer](
		"transfer_file",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, transfer core.FileTransfer) error {
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
		func(_ context.Context) ([]core.FileTransfer, error) {
			return useCases.FileTransferList(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

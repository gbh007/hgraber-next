package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/agentmodel"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type exportUnitUseCases interface {
	ExportList() []agentmodel.BookToExport
	ExportArchive(ctx context.Context, book agentmodel.BookToExport, retry bool) error
}

func NewExporter(
	useCases exportUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[agentmodel.BookToExport] {
	return worker.New[agentmodel.BookToExport](
		cfg,
		logger,
		func(ctx context.Context, toExport agentmodel.BookToExport) error {
			err := useCases.ExportArchive(ctx, toExport, true)
			if err != nil {
				return pkg.WrapError(
					err, "fail export book",
					pkg.ErrorArgument("book_id", toExport.BookID),
					pkg.ErrorArgument("agent_id", toExport.AgentID),
				)
			}

			return nil
		},
		func(_ context.Context) ([]agentmodel.BookToExport, error) {
			return useCases.ExportList(), nil
		},
		tracer,
		metricProvider,
	)
}

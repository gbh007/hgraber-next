package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type exportUnitUseCases interface {
	ExportList() []serveragentdto.BookToExport
	ExportArchive(ctx context.Context, book serveragentdto.BookToExport, retry bool) error
}

func NewExporter(
	useCases exportUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[serveragentdto.BookToExport] {
	return worker.New[serveragentdto.BookToExport](
		cfg,
		logger,
		func(ctx context.Context, toExport serveragentdto.BookToExport) error {
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
		func(_ context.Context) ([]serveragentdto.BookToExport, error) {
			return useCases.ExportList(), nil
		},
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type exportUnitUseCases interface {
	ExportList() []core.BookFullWithAgent
	ExportArchive(ctx context.Context, book core.BookFullWithAgent, retry bool) error
}

func NewExporter(
	useCases exportUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[core.BookFullWithAgent] {
	return worker.New[core.BookFullWithAgent](
		"export",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, book core.BookFullWithAgent) error {
			err := useCases.ExportArchive(ctx, book, true)
			if err != nil {
				return pkg.WrapError(
					err, "fail export book",
					pkg.ErrorArgument("book_id", book.Book.ID),
					pkg.ErrorArgument("agent_id", book.AgentID),
				)
			}

			return nil
		},
		func(_ context.Context) ([]core.BookFullWithAgent, error) {
			return useCases.ExportList(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

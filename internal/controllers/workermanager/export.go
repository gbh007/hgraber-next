package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type exportUnitUseCases interface {
	ExportList() []entities.BookFullWithAgent
	ExportArchive(ctx context.Context, book entities.BookFullWithAgent, retry bool) error
}

func NewExporter(
	useCases exportUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[entities.BookFullWithAgent] {
	return worker.New[entities.BookFullWithAgent](
		"export",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, book entities.BookFullWithAgent) error {
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
		func(_ context.Context) ([]entities.BookFullWithAgent, error) {
			return useCases.ExportList(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

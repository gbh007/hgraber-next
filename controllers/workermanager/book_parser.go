package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type bookWorkerUnitUseCases interface {
	ParseBook(ctx context.Context, agentID uuid.UUID, book core.Book) error
	BooksToParse(ctx context.Context) ([]core.BookWithAgent, error)
}

func NewBookParser(
	useCases bookWorkerUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[core.BookWithAgent] {
	return worker.New[core.BookWithAgent](
		"book",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, book core.BookWithAgent) error {
			err := useCases.ParseBook(ctx, book.AgentID, book.Book)
			if err != nil {
				return pkg.WrapError(
					err, "fail parse book",
					pkg.ErrorArgument("book_id", book.ID),
				)
			}

			return nil
		},
		useCases.BooksToParse,
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
)

type bookWorkerUnitUseCases interface {
	ParseBook(ctx context.Context, agentID uuid.UUID, book entities.Book) error
	BooksToParse(ctx context.Context) ([]entities.BookWithAgent, error)
}

func NewBookParser(
	useCases bookWorkerUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg WorkerConfig,
) *worker.Worker[entities.BookWithAgent] {
	return worker.New[entities.BookWithAgent](
		"book",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, book entities.BookWithAgent) {
			err := useCases.ParseBook(ctx, book.AgentID, book.Book)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail parse book",
					slog.String("book_id", book.ID.String()),
					slog.Any("error", err),
				)
			}
		},
		func(ctx context.Context) []entities.BookWithAgent {
			books, err := useCases.BooksToParse(ctx)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail get books for parse",
					slog.Any("error", err),
				)
			}

			return books
		},
		cfg.GetCount(),
		tracer,
	)
}

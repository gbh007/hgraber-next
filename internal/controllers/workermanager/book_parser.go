package workermanager

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
)

type bookWorkerUnitUseCases interface {
	ParseBook(ctx context.Context, agentID uuid.UUID, book entities.Book) error
	BooksToParse(ctx context.Context) ([]entities.BookWithAgent, error)
}

func NewBookParser(useCases bookWorkerUnitUseCases, logger *slog.Logger) *worker.Worker[entities.BookWithAgent] {
	return worker.New[entities.BookWithAgent](
		"book",
		1000,
		time.Second*15,
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
		10,
	)
}

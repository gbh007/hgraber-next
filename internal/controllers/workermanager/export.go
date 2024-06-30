package workermanager

import (
	"context"
	"log/slog"
	"time"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
)

type exportUnitUseCases interface {
	ExportList() []entities.BookFullWithAgent
	ExportArchive(ctx context.Context, book entities.BookFullWithAgent, retry bool) error
}

func NewExporter(useCases exportUnitUseCases, logger *slog.Logger) *worker.Worker[entities.BookFullWithAgent] {
	return worker.New[entities.BookFullWithAgent](
		"export",
		1000,
		time.Minute,
		logger,
		func(ctx context.Context, book entities.BookFullWithAgent) {
			err := useCases.ExportArchive(ctx, book, true)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail export book",
					slog.String("book_id", book.ID.String()),
					slog.String("agent_id", book.AgentID.String()),
					slog.Any("error", err),
				)
			}
		},
		func(_ context.Context) []entities.BookFullWithAgent {
			return useCases.ExportList()
		},
		3,
	)
}
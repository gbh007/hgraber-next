package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

type bookWorkerUnitUseCases interface {
	ParseBook(ctx context.Context, agentID uuid.UUID, book core.Book) error
	BooksToParse(ctx context.Context) ([]serveragentdto.BookWithAgent, error)
}

func NewBookParser(
	useCases bookWorkerUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[serveragentdto.BookWithAgent] {
	return worker.New[serveragentdto.BookWithAgent](
		cfg,
		logger,
		func(ctx context.Context, book serveragentdto.BookWithAgent) error {
			err := useCases.ParseBook(ctx, book.AgentID, book.Book)
			if err != nil {
				return fmt.Errorf("fail parse book (book_id: %s): %w", book.ID.String(), err)
			}

			return nil
		},
		useCases.BooksToParse,
		metricProvider,
	)
}

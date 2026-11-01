package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

type bookWorkerUnitUseCases interface {
	ParseBook(ctx context.Context, agentID uuid.UUID, book core.Book) error
	BooksToParse(ctx context.Context) ([]serveragentdto.BookWithAgent, error)
}

func NewBookParser(
	useCases bookWorkerUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[serveragentdto.BookWithAgent] {
	return worker.New[serveragentdto.BookWithAgent](
		cfg,
		logger,
		func(ctx context.Context, book serveragentdto.BookWithAgent) error {
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
		tracer,
		metricProvider,
	)
}

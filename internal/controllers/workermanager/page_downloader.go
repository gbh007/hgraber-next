package workermanager

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
)

type pageWorkerUnitUseCases interface {
	DownloadPage(ctx context.Context, agentID uuid.UUID, page entities.PageForDownload) error
	PagesToDownload(ctx context.Context) ([]entities.PageForDownloadWithAgent, error)
}

func NewPageDownloader(useCases pageWorkerUnitUseCases, logger *slog.Logger) *worker.Worker[entities.PageForDownloadWithAgent] {
	return worker.New[entities.PageForDownloadWithAgent](
		"page",
		10000,
		time.Second*15,
		logger,
		func(ctx context.Context, page entities.PageForDownloadWithAgent) {
			err := useCases.DownloadPage(ctx, page.AgentID, page.PageForDownload)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail download page",
					slog.String("book_id", page.BookID.String()),
					slog.Int("page_number", page.PageNumber),
					slog.Any("error", err),
				)
			}
		},
		func(ctx context.Context) []entities.PageForDownloadWithAgent {
			pages, err := useCases.PagesToDownload(ctx)
			if err != nil {
				logger.ErrorContext(
					ctx, "fail get pages for download",
					slog.Any("error", err),
				)
			}

			return pages
		},
		10,
	)
}

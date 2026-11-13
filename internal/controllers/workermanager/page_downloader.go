package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
)

type pageWorkerUnitUseCases interface {
	DownloadPage(ctx context.Context, agentID uuid.UUID, page core.PageForDownload) error
	PagesToDownload(ctx context.Context) ([]parsing.PageForDownloadWithAgent, error)
}

func NewPageDownloader(
	useCases pageWorkerUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[parsing.PageForDownloadWithAgent] {
	return worker.New[parsing.PageForDownloadWithAgent](
		cfg,
		logger,
		func(ctx context.Context, page parsing.PageForDownloadWithAgent) error {
			err := useCases.DownloadPage(ctx, page.AgentID, page.PageForDownload)
			if err != nil {
				return fmt.Errorf(
					"fail download page (book_id: %s, page_number: %d): %w",
					page.BookID.String(), page.PageNumber, err,
				)
			}

			return nil
		},
		useCases.PagesToDownload,
		metricProvider,
	)
}

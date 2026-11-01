package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/pkg"
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
				return pkg.WrapError(
					err, "fail download page",
					pkg.ErrorArgument("book_id", page.BookID),
					pkg.ErrorArgument("page_number", page.PageNumber),
				)
			}

			return nil
		},
		useCases.PagesToDownload,
		metricProvider,
	)
}

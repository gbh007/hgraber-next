package workermanager

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/pkg"
)

type pageWorkerUnitUseCases interface {
	DownloadPage(ctx context.Context, agentID uuid.UUID, page entities.PageForDownload) error
	PagesToDownload(ctx context.Context) ([]entities.PageForDownloadWithAgent, error)
}

func NewPageDownloader(
	useCases pageWorkerUnitUseCases,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[entities.PageForDownloadWithAgent] {
	return worker.New[entities.PageForDownloadWithAgent](
		"page",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, page entities.PageForDownloadWithAgent) error {
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
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/internal/controllers/workermanager/internal/worker"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

type exportUnitUseCases interface {
	ExportList() []serveragentdto.BookToExport
	ExportArchive(ctx context.Context, book serveragentdto.BookToExport, retry bool) error
}

func NewExporter(
	useCases exportUnitUseCases,
	logger *slog.Logger,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[serveragentdto.BookToExport] {
	return worker.New[serveragentdto.BookToExport](
		cfg,
		logger,
		func(ctx context.Context, toExport serveragentdto.BookToExport) error {
			err := useCases.ExportArchive(ctx, toExport, true)
			if err != nil {
				return fmt.Errorf(
					"fail export book (book_id: %s, agent_id: %s): %w",
					toExport.BookID.String(), toExport.AgentID.String(), err,
				)
			}

			return nil
		},
		func(_ context.Context) ([]serveragentdto.BookToExport, error) {
			return useCases.ExportList(), nil
		},
		metricProvider,
	)
}

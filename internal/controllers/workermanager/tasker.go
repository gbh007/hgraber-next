package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/internal/entities"
)

type taskerDataProvider interface {
	GetTask() []entities.RunnableTask
	SaveTaskResult(result *entities.TaskResult)
}

func NewTasker(
	useCases taskerDataProvider,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[entities.RunnableTask] {
	return worker.New[entities.RunnableTask](
		"task",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, task entities.RunnableTask) error {
			tr := new(entities.TaskResult)
			useCases.SaveTaskResult(tr)

			task.Run(ctx, tr)

			return nil
		},
		func(ctx context.Context) ([]entities.RunnableTask, error) {
			return useCases.GetTask(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

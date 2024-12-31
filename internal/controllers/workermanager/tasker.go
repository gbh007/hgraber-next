package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/controllers/internal/worker"
	"hgnext/internal/entities"
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
		func(ctx context.Context, task entities.RunnableTask) {
			tr := new(entities.TaskResult)
			useCases.SaveTaskResult(tr)

			task.Run(ctx, tr)
		},
		func(ctx context.Context) []entities.RunnableTask {
			return useCases.GetTask()
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

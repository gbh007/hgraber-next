package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/core"
)

type taskerDataProvider interface {
	GetTask() []core.RunnableTask
	SaveTaskResult(result *core.TaskResult)
}

func NewTasker(
	useCases taskerDataProvider,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[core.RunnableTask] {
	return worker.New[core.RunnableTask](
		"task",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, task core.RunnableTask) error {
			tr := new(core.TaskResult)
			useCases.SaveTaskResult(tr)

			task.Run(ctx, tr)

			return nil
		},
		func(ctx context.Context) ([]core.RunnableTask, error) {
			return useCases.GetTask(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/internal/worker"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type taskerDataProvider interface {
	GetTask() []systemmodel.RunnableTask
	SaveTaskResult(result *systemmodel.TaskResult)
}

func NewTasker(
	useCases taskerDataProvider,
	logger *slog.Logger,
	tracer trace.Tracer,
	cfg workerConfig,
	metricProvider metricProvider,
) *worker.Worker[systemmodel.RunnableTask] {
	return worker.New[systemmodel.RunnableTask](
		"task",
		cfg.GetQueueSize(),
		cfg.GetInterval(),
		logger,
		func(ctx context.Context, task systemmodel.RunnableTask) error {
			tr := new(systemmodel.TaskResult)
			useCases.SaveTaskResult(tr)

			task.Run(ctx, tr)

			return nil
		},
		func(ctx context.Context) ([]systemmodel.RunnableTask, error) {
			return useCases.GetTask(), nil
		},
		cfg.GetCount(),
		tracer,
		metricProvider,
	)
}

package workermanager

import (
	"context"
	"log/slog"
	"sync"

	"hgnext/internal/entities"
)

type WorkerUnit interface {
	Name() string
	Serve(ctx context.Context)

	InQueueCount() int
	InWorkCount() int
	RunnersCount() int
}

type Controller struct {
	workerUnits []WorkerUnit
	logger      *slog.Logger
}

func New(logger *slog.Logger, workerUnits ...WorkerUnit) *Controller {
	return &Controller{
		logger:      logger,
		workerUnits: workerUnits,
	}
}

func (c *Controller) Name() string {
	return "worker manager"
}

func (c *Controller) Start(ctx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	wg := new(sync.WaitGroup)

	wg.Add(len(c.workerUnits))

	for _, w := range c.workerUnits {
		go func(ctx context.Context, w WorkerUnit) {
			defer wg.Done()

			w.Serve(ctx)
		}(ctx, w)
	}

	go func() {
		defer close(done)

		c.logger.InfoContext(ctx, "worker manager start")
		defer c.logger.InfoContext(ctx, "worker manager stop")

		wg.Wait()
	}()

	return done, nil
}

func (c *Controller) Info() []entities.SystemWorkerStat {
	res := make([]entities.SystemWorkerStat, 0, len(c.workerUnits))

	for _, worker := range c.workerUnits {
		res = append(res, entities.SystemWorkerStat{
			Name:         worker.Name(),
			InQueueCount: worker.InQueueCount(),
			InWorkCount:  worker.InWorkCount(),
			RunnersCount: worker.RunnersCount(),
		})
	}

	return res
}

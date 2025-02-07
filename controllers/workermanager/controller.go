package workermanager

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type WorkerUnit interface {
	Name() string
	Serve(ctx context.Context)

	InQueueCount() int
	InWorkCount() int
	RunnersCount() int

	SetRunnersCount(newUnitCount int)
}

type workerConfig interface {
	GetCount() int32
	GetQueueSize() int
	GetInterval() time.Duration
}

type metricProvider interface {
	RegisterWorkerExecutionTaskTime(name string, d time.Duration, success bool)
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

func (c *Controller) Info() []systemmodel.SystemWorkerStat {
	res := make([]systemmodel.SystemWorkerStat, 0, len(c.workerUnits))

	for _, worker := range c.workerUnits {
		res = append(res, systemmodel.SystemWorkerStat{
			Name:         worker.Name(),
			InQueueCount: worker.InQueueCount(),
			InWorkCount:  worker.InWorkCount(),
			RunnersCount: worker.RunnersCount(),
		})
	}

	return res
}

func (c *Controller) SetRunnerCount(_ context.Context, counts map[string]int) {
	for _, worker := range c.workerUnits {
		count, ok := counts[worker.Name()]
		if ok {
			worker.SetRunnersCount(count)
		}
	}
}

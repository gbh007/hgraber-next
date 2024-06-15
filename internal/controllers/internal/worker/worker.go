package worker

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"
)

type Worker[T any] struct {
	name  string
	queue chan T

	inWorkRunnersCount *atomic.Int32
	runnersCount       *atomic.Int32

	interval time.Duration

	handler func(context.Context, T)
	getter  func(context.Context) []T

	logger *slog.Logger
}

func New[T any](
	name string,
	queueSize int,
	interval time.Duration,
	logger *slog.Logger,
	handler func(context.Context, T),
	getter func(context.Context) []T,
	runnersCount int32,
) *Worker[T] {
	w := &Worker[T]{
		name:               name,
		queue:              make(chan T, queueSize),
		inWorkRunnersCount: new(atomic.Int32),
		runnersCount:       new(atomic.Int32),
		interval:           interval,
		handler:            handler,
		getter:             getter,

		logger: logger,
	}

	w.runnersCount.Store(runnersCount)

	return w
}

func (w *Worker[T]) InQueueCount() int {
	return len(w.queue)
}

func (w *Worker[T]) InWorkCount() int {
	return int(w.inWorkRunnersCount.Load())
}

func (w *Worker[T]) RunnersCount() int {
	return int(w.runnersCount.Load())
}

func (w *Worker[T]) handleOne(ctx context.Context, value T) {
	defer func() {
		if p := recover(); p != nil {
			w.logger.WarnContext(
				ctx, "panic in worker detected",
				slog.Any("panic", p),
				slog.String("worker_name", w.name),
			)
		}
	}()

	w.inWorkRunnersCount.Add(1)
	defer w.inWorkRunnersCount.Add(-1)

	w.handler(ctx, value)
}

func (w *Worker[T]) runQueueHandler(ctx context.Context) {
	w.logger.DebugContext(ctx, "worker handler start", slog.String("worker_name", w.name))
	defer w.logger.DebugContext(ctx, "worker handler stop", slog.String("worker_name", w.name))

	for {
		select {
		case value := <-w.queue:
			w.handleOne(ctx, value)
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker[T]) Serve(ctx context.Context) {
	wg := new(sync.WaitGroup)

	handlersCount := int(w.runnersCount.Load())

	for i := 0; i < handlersCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			w.runQueueHandler(ctx)
		}()
	}

	w.logger.DebugContext(ctx, "worker start", slog.String("worker_name", w.name))
	defer w.logger.DebugContext(ctx, "worker stop", slog.String("worker_name", w.name))

	timer := time.NewTicker(w.interval)

handler:
	for {
		select {
		case <-ctx.Done():
			break handler

		case <-timer.C:
			if len(w.queue) > 0 {
				continue
			}

			for _, title := range w.getter(ctx) {
				select {
				case <-ctx.Done():
					break handler

				case w.queue <- title:
				}

			}
		}
	}

	// Дожидаемся завершения всех подпроцессов
	wg.Wait()
}

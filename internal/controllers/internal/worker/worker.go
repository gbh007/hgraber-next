package worker

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/trace"
)

type metricProvider interface {
	RegisterWorkerExecutionTaskTime(name string, d time.Duration)
}

type Worker[T any] struct {
	name  string
	queue chan T

	inWorkRunnersCount  atomic.Int32
	runnersCount        atomic.Int32
	defaultRunnersCount int

	interval time.Duration

	handler func(context.Context, T)
	getter  func(context.Context) []T

	logger         *slog.Logger
	tracer         trace.Tracer
	metricProvider metricProvider

	unitsMutex sync.Mutex
	units      []*Unit[T]
	unitsWG    sync.WaitGroup
	unitCtx    context.Context
}

func New[T any](
	name string,
	queueSize int,
	interval time.Duration,
	logger *slog.Logger,
	handler func(context.Context, T),
	getter func(context.Context) []T,
	runnersCount int32,
	tracer trace.Tracer,
	metricProvider metricProvider,
) *Worker[T] {
	w := &Worker[T]{
		name:     name,
		queue:    make(chan T, queueSize),
		interval: interval,
		handler:  handler,
		getter:   getter,

		logger:         logger,
		tracer:         tracer,
		metricProvider: metricProvider,

		defaultRunnersCount: int(runnersCount),
	}

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

func (w *Worker[T]) Name() string {
	return w.name
}

func (w *Worker[T]) SetRunnersCount(newUnitCount int) {
	go func() {
		w.unitsMutex.Lock()
		defer w.unitsMutex.Unlock()

		if newUnitCount < 0 {
			newUnitCount = 0
		}

		oldUnitCount := len(w.units)

		if oldUnitCount < newUnitCount {
			for i := oldUnitCount; i < newUnitCount; i++ {
				unit := NewUnit(
					w.name,
					int32(i),
					w.logger,
					w.handler,
					w.tracer,
					w.metricProvider,
					w.queue,
					UnitCallback{
						StartHandleOne:  func() { w.inWorkRunnersCount.Add(1) },
						FinishHandleOne: func() { w.inWorkRunnersCount.Add(-1) },
						StartUnit: func() {
							w.runnersCount.Add(1)
							w.unitsWG.Add(1)
						},
						StopUnit: func() {
							w.runnersCount.Add(-1)
							w.unitsWG.Done()
						},
					},
				)

				w.units = append(w.units, unit)

				go unit.Serve(w.unitCtx)
			}
		}

		if oldUnitCount > newUnitCount {
			for i := newUnitCount; i < oldUnitCount; i++ {
				w.units[i].ShutDown(context.Background())
			}

			w.units = w.units[:newUnitCount]
		}
	}()
}

func (w *Worker[T]) Serve(ctx context.Context) {
	w.logger.DebugContext(ctx, "worker start", slog.String("worker_name", w.name))
	defer w.logger.DebugContext(ctx, "worker stop", slog.String("worker_name", w.name))

	w.unitCtx = ctx

	w.SetRunnersCount(w.defaultRunnersCount)

	timer := time.NewTicker(w.interval)

handler:
	for {
		select {
		case <-ctx.Done():
			break handler

		case <-timer.C:
			// FIXME: сейчас это скорее заглушка, чтобы не было избыточных переобработок.
			if len(w.queue) > 0 || w.inWorkRunnersCount.Load() > 0 {
				continue
			}

			ctx, span := w.tracer.Start(
				ctx, "worker-fetch/"+w.name,
				trace.WithSpanKind(trace.SpanKindServer),
			)

			for _, data := range w.getter(ctx) {
				select {
				case <-ctx.Done():
					span.End()
					break handler

				case w.queue <- data:
				}
			}

			span.End()
		}
	}

	// Дожидаемся завершения всех подпроцессов
	w.unitsWG.Wait()
}

package worker

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type UnitCallback struct {
	StartHandleOne  func()
	FinishHandleOne func()

	StartUnit func()
	StopUnit  func()
}

type Unit[T any] struct {
	name   string
	number int32

	queue <-chan T

	handler func(context.Context, T)

	logger         *slog.Logger
	tracer         trace.Tracer
	metricProvider metricProvider

	callback UnitCallback

	cancel context.CancelFunc
	wait   sync.WaitGroup
}

func NewUnit[T any](
	name string,
	number int32,
	logger *slog.Logger,
	handler func(context.Context, T),
	tracer trace.Tracer,
	metricProvider metricProvider,
	queue <-chan T,
	callback UnitCallback,
) *Unit[T] {
	w := &Unit[T]{
		name:    name,
		queue:   queue,
		handler: handler,

		logger:         logger,
		tracer:         tracer,
		metricProvider: metricProvider,

		callback: callback,
		number:   number,

		cancel: func() {},
	}

	return w
}

func (w *Unit[T]) Name() string {
	return w.name
}

func (w *Unit[T]) handleOne(ctx context.Context, value T) {
	defer func() {
		if p := recover(); p != nil {
			w.logger.WarnContext(
				ctx, "panic in worker unit detected",
				slog.Any("panic", p),
				slog.String("worker_name", w.name),
				slog.Int("worker_unit", int(w.number)),
			)
		}
	}()

	ctx, span := w.tracer.Start(
		ctx, "worker-job/"+w.name,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	span.SetAttributes(attribute.Int("hgnext.worker.unit", int(w.number)))

	w.callback.StartHandleOne()
	defer w.callback.FinishHandleOne()

	tStart := time.Now()
	defer func() {
		w.metricProvider.RegisterWorkerExecutionTaskTime(w.name, time.Since(tStart))
	}()

	ctx = context.WithoutCancel(ctx)

	w.handler(ctx, value)
}

func (w *Unit[T]) Serve(ctx context.Context) {
	ctx, w.cancel = context.WithCancel(ctx)

	w.wait.Add(1)
	defer w.wait.Done()

	w.logger.DebugContext(
		ctx, "worker unit start",
		slog.String("worker_name", w.name),
		slog.Int("worker_unit", int(w.number)),
	)
	defer w.logger.DebugContext(
		ctx, "worker unit stop",
		slog.String("worker_name", w.name),
		slog.Int("worker_unit", int(w.number)),
	)

	w.callback.StartUnit()
	defer w.callback.StopUnit()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case value := <-w.queue:
			w.handleOne(ctx, value)
		case <-ctx.Done():
			return
		}
	}
}

func (w *Unit[T]) ShutDown(_ context.Context) {
	w.cancel()
	w.wait.Wait()
}

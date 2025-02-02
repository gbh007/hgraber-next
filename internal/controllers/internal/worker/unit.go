package worker

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strconv"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type workerHandlerFunc[T any] func(context.Context, T) error

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

	handler workerHandlerFunc[T]

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
	handler workerHandlerFunc[T],
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

func (w *Unit[T]) handleOne(ctx context.Context, value T) (err error) {
	tStart := time.Now()
	defer func() {
		w.metricProvider.RegisterWorkerExecutionTaskTime(w.name, time.Since(tStart), err == nil)
	}()

	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic detected: %v", p)
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

	ctx = context.WithoutCancel(ctx)

	return w.handler(ctx, value)
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
			err := w.handleOne(ctx, value)
			if err != nil {
				w.logger.ErrorContext(
					ctx, "worker fail task",
					slog.String("worker_name", w.name),
					slog.Int("worker_unit", int(w.number)),
					slog.Any("error", err),
				)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (w *Unit[T]) ShutDown(_ context.Context) {
	w.cancel()
	w.wait.Wait()
}

func stackTrace(skip, count int) []string {
	result := []string{}

	pc := make([]uintptr, count)
	n := runtime.Callers(skip, pc)

	pc = pc[:n]

	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()

		result = append(result, frame.File+":"+strconv.Itoa(frame.Line))

		if !more {
			break
		}
	}

	return result
}

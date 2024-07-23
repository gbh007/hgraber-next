package logAdapter

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type Adapter struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Adapter {
	return &Adapter{
		logger: logger,
	}
}

func (a *Adapter) Logger(ctx context.Context) *slog.Logger {
	snapContext := trace.SpanContextFromContext(ctx)
	if !snapContext.HasTraceID() {
		return a.logger
	}

	return a.logger.With(slog.String("trace_id", snapContext.TraceID().String()))
}

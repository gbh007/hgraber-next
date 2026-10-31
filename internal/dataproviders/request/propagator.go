package request

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
)

type noopPropagator struct{}

func (noopPropagator) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {}
func (noopPropagator) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	return ctx
}
func (noopPropagator) Fields() []string { return nil }

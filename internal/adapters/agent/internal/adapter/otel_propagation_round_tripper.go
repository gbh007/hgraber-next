package adapter

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type otelPropagationRoundTripper struct {
	next http.RoundTripper
}

func (rt otelPropagationRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	otel.GetTextMapPropagator().Inject( //nolint:contextcheck // ложно-положительное срабатывание
		req.Context(),
		propagation.HeaderCarrier(req.Header),
	)

	return rt.next.RoundTrip(req) //nolint:wrapcheck // специально не вмешиваемся в работу
}

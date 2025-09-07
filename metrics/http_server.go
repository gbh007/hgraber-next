package metrics

import (
	"time"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpServerHandleRequest = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metriccore.SystemName,
			Subsystem: "http_server",
			Name:      "handle_duration",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.2, 0.5, 1, 1.5, 2, 3, 5, 10},
			Help:      "Время обработки запроса",
		},
		[]string{"server_addr", "operation", "status"},
	)
	httpServerActiveRequest = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metriccore.SystemName,
			Subsystem: "http_server",
			Name:      "active_handlers_total",
			Help:      "Количество активных запросов",
		},
		[]string{"server_addr", "operation"},
	)
)

func (MetricProvider) HTTPServerAddHandle(addr, operation string, status bool, d time.Duration) {
	httpServerHandleRequest.WithLabelValues(addr, operation, metriccore.ErrorLabel(status)).Observe(d.Seconds())
}

func (MetricProvider) HTTPServerIncActive(addr, operation string) {
	httpServerActiveRequest.WithLabelValues(addr, operation).Inc()
}

func (MetricProvider) HTTPServerDecActive(addr, operation string) {
	httpServerActiveRequest.WithLabelValues(addr, operation).Dec()
}

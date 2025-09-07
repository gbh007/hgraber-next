package metrics

import (
	"time"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	dbActiveRequest = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metriccore.SystemName,
		Subsystem: "database",
		Name:      "active_request",
		Help:      "Количество активных запросов",
	})
	dbOpenConnection = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metriccore.SystemName,
		Subsystem: "database",
		Name:      "open_connection",
		Help:      "Количество открытых соединений",
	})
	dbRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metriccore.SystemName,
			Subsystem: "database",
			Name:      "request_duration",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.2, 0.5, 1, 1.5, 2, 3, 5, 10},
			Help:      "Время обработки запроса",
		},
		[]string{"stmt"},
	)
)

func (MetricProvider) IncDBActiveRequest() {
	dbActiveRequest.Inc()
}

func (MetricProvider) DecDBActiveRequest() {
	dbActiveRequest.Dec()
}

func (MetricProvider) SetDBOpenConnection(n int32) {
	dbOpenConnection.Set(float64(n))
}

func (MetricProvider) DecDBOpenConnection() {
	dbOpenConnection.Dec()
}

func (MetricProvider) RegisterDBRequestDuration(stmt string, d time.Duration) {
	dbRequestDuration.WithLabelValues(stmt).Observe(d.Seconds())
}

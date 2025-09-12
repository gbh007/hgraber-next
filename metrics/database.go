package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/gbh007/hgraber-next/metrics/metricdatabase"
)

var (
	dbActiveRequest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricdatabase.ActiveRequestName,
		Help: "Количество активных запросов",
	})
	dbOpenConnection = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricdatabase.OpenConnectionName,
		Help: "Количество открытых соединений",
	})
	dbRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricdatabase.RequestDurationName,
			Buckets: []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.2, 0.5, 1, 1.5, 2, 3, 5, 10},
			Help:    "Время обработки запроса",
		},
		[]string{metricdatabase.StmtLabelName},
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

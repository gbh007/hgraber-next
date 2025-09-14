package metric

import (
	"time"

	"github.com/gbh007/hgraber-next/adapters/metric/metricdatabase"
)

func (MetricProvider) IncDBActiveRequest(db string) {
	metricdatabase.ActiveRequest.WithLabelValues(db).Inc()
}

func (MetricProvider) DecDBActiveRequest(db string) {
	metricdatabase.ActiveRequest.WithLabelValues(db).Dec()
}

func (MetricProvider) SetDBOpenConnection(db string, n int32) {
	metricdatabase.OpenConnection.WithLabelValues(db).Set(float64(n))
}

func (MetricProvider) DecDBOpenConnection(db string) {
	metricdatabase.OpenConnection.WithLabelValues(db).Dec()
}

func (MetricProvider) RegisterDBRequestDuration(db, stmt string, d time.Duration) {
	metricdatabase.RequestDuration.WithLabelValues(db, stmt).Observe(d.Seconds())
}

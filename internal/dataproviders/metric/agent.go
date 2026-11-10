package metric

import (
	"time"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricagent"
)

func (MetricProvider) RegisterParserActionTime(action, parserName string, d time.Duration) {
	metricagent.ParserActionTime.WithLabelValues(action, parserName).Observe(d.Seconds())
}

func (MetricProvider) IncWebCacheCounter(action string) {
	metricagent.WebCacheCounter.WithLabelValues(action).Inc()
}

func (MetricProvider) SetWebCacheSize(size int64) {
	metricagent.WebCacheBytes.Set(float64(size))
}

func (MetricProvider) SetWebCacheFileTotal(count int64) {
	metricagent.WebCacheFileTotal.Set(float64(count))
}

package metric

import (
	"time"

	"github.com/gbh007/hgraber-next/adapters/metric/metricagent"
)

func (MetricProvider) RegisterParserActionTime(action, parserName string, d time.Duration) {
	metricagent.ParserActionTime.WithLabelValues(action, parserName).Observe(d.Seconds())
}

func (MetricProvider) IncWebCacheCounter(action string) {
	metricagent.WebCacheCounter.WithLabelValues(action).Inc()
}

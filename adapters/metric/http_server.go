package metric

import (
	"time"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metrichttp"
)

func (MetricProvider) HTTPServerAddHandle(addr, operation string, status bool, d time.Duration) {
	metrichttp.ServerHandleRequest.WithLabelValues(addr, operation, metriccore.ErrorLabel(status)).Observe(d.Seconds())
}

func (MetricProvider) HTTPServerIncActive(addr, operation string) {
	metrichttp.ServerActiveRequest.WithLabelValues(addr, operation).Inc()
}

func (MetricProvider) HTTPServerDecActive(addr, operation string) {
	metrichttp.ServerActiveRequest.WithLabelValues(addr, operation).Dec()
}

package metric

import (
	"time"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func (MetricProvider) RegisterWorkerExecutionTaskTime(name string, d time.Duration, success bool) {
	metricserver.WorkerExecutionTaskTime.WithLabelValues(name, metriccore.ErrorLabel(success)).Observe(d.Seconds())
}

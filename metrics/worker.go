package metrics

import (
	"time"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var workerExecutionTaskTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace: metriccore.SystemName,
	Subsystem: metriccore.SubSystemName,
	Name:      "worker_execution_task_seconds",
	Help:      "Время выполнения задачи воркером",
}, []string{"worker_name", "success"})

func (MetricProvider) RegisterWorkerExecutionTaskTime(name string, d time.Duration, success bool) {
	workerExecutionTaskTime.WithLabelValues(name, metriccore.ErrorLabel(success)).Observe(d.Seconds())
}

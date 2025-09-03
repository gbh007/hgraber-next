package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var workerExecutionTaskTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace: SystemName,
	Subsystem: SubSystemName,
	Name:      "worker_execution_task_seconds",
	Help:      "Время выполнения задачи воркером",
}, []string{"worker_name", "success"})

func (MetricProvider) RegisterWorkerExecutionTaskTime(name string, d time.Duration, success bool) {
	workerExecutionTaskTime.WithLabelValues(name, errorLabel(success)).Observe(d.Seconds())
}

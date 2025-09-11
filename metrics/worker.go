package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

var workerExecutionTaskTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: metricserver.WorkerExecutionTaskSecondsName,
	Help: "Время выполнения задачи воркером",
}, []string{metricserver.WorkerNameLabel, metriccore.SuccessLabel})

func (MetricProvider) RegisterWorkerExecutionTaskTime(name string, d time.Duration, success bool) {
	workerExecutionTaskTime.WithLabelValues(name, metriccore.ErrorLabel(success)).Observe(d.Seconds())
}

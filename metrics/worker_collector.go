package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

type workerInfoProvider interface {
	WorkersInfo(ctx context.Context) []systemmodel.SystemWorkerStat
}

var workerDesc = prometheus.NewDesc(
	metricserver.WorkerTotalName,
	"Данные воркеров",
	[]string{metricserver.WorkerNameLabel, metricserver.CounterLabel}, nil,
)

var _ prometheus.Collector = (*WorkerInfoCollector)(nil)

type WorkerInfoCollector struct {
	logger             *slog.Logger
	workerInfoProvider workerInfoProvider

	timeout time.Duration
}

func RegisterWorkerInfoCollector(
	logger *slog.Logger,
	workerInfoProvider workerInfoProvider,
	timeout time.Duration,
) error {
	return prometheus.Register(&WorkerInfoCollector{
		logger:             logger,
		workerInfoProvider: workerInfoProvider,
		timeout:            timeout,
	})
}

func (c *WorkerInfoCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- workerDesc
}

func (c *WorkerInfoCollector) Collect(metr chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	workers := c.workerInfoProvider.WorkersInfo(ctx)

	for _, worker := range workers {
		metr <- prometheus.MustNewConstMetric(workerDesc, prometheus.GaugeValue, float64(worker.InQueueCount), worker.Name, "in_queue")

		metr <- prometheus.MustNewConstMetric(workerDesc, prometheus.GaugeValue, float64(worker.InWorkCount), worker.Name, "in_work")

		metr <- prometheus.MustNewConstMetric(workerDesc, prometheus.GaugeValue, float64(worker.RunnersCount), worker.Name, "runners")
	}
}

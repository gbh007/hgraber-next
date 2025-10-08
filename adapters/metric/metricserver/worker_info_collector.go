package metricserver

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

var _ prometheus.Collector = (*WorkerInfoCollector)(nil)

type workerInfoProvider interface {
	WorkersInfo(ctx context.Context) []systemmodel.SystemWorkerStat
}

type WorkerInfoCollector struct {
	workerInfoProvider workerInfoProvider
}

func RegisterWorkerInfoCollector(
	registerer prometheus.Registerer,
	workerInfoProvider workerInfoProvider,
) error {
	return registerer.Register(&WorkerInfoCollector{ //nolint:wrapcheck // не имеет смысла
		workerInfoProvider: workerInfoProvider,
	})
}

func (c *WorkerInfoCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- WorkerDesc
}

func (c *WorkerInfoCollector) Collect(metr chan<- prometheus.Metric) {
	ctx := context.Background()

	workers := c.workerInfoProvider.WorkersInfo(ctx)

	for _, worker := range workers {
		metr <- prometheus.MustNewConstMetric(
			WorkerDesc,
			prometheus.GaugeValue,
			float64(worker.InQueueCount),
			worker.Name,
			"in_queue",
		)

		metr <- prometheus.MustNewConstMetric(
			WorkerDesc,
			prometheus.GaugeValue,
			float64(worker.InWorkCount),
			worker.Name,
			"in_work",
		)

		metr <- prometheus.MustNewConstMetric(
			WorkerDesc,
			prometheus.GaugeValue,
			float64(worker.RunnersCount),
			worker.Name,
			"runners",
		)
	}
}

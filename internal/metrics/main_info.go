package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"hgnext/internal/entities"
)

type infoProvider interface {
	SystemInfo(ctx context.Context) (entities.SystemSizeInfoWithMonitor, error)
}

var (
	bookDesc = prometheus.NewDesc(
		"book_total",
		"Количество книг по статусам",
		[]string{"type"}, nil,
	)
	pageDesc = prometheus.NewDesc(
		"page_total",
		"Количество страниц по статусам",
		[]string{"type"}, nil,
	)
	fileDesc = prometheus.NewDesc(
		"file_bytes",
		"Размер файлов по статусам",
		[]string{"type"}, nil,
	)
	workerDesc = prometheus.NewDesc(
		"worker_total",
		"Данные воркеров",
		[]string{"worker_name", "counter"}, nil,
	)
)

var _ prometheus.Collector = (*SystemInfoCollector)(nil)

type SystemInfoCollector struct {
	logger       *slog.Logger
	infoProvider infoProvider
}

func RegisterSystemInfoCollector(
	logger *slog.Logger,
	infoProvider infoProvider,
) error {
	return prometheus.Register(&SystemInfoCollector{
		logger:       logger,
		infoProvider: infoProvider,
	})
}

func (c *SystemInfoCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- bookDesc
	desc <- pageDesc
	desc <- fileDesc
}

//nolint:promlinter // ложно-положительное срабатывание
func (c *SystemInfoCollector) Collect(metr chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	res, err := c.infoProvider.SystemInfo(ctx)
	if err != nil {
		c.logger.ErrorContext(
			ctx, "failed scrap system info",
			slog.Any("error", err),
		)

		return
	}

	metr <- prometheus.MustNewConstMetric(bookDesc, prometheus.GaugeValue, float64(res.BookCount), "all")
	metr <- prometheus.MustNewConstMetric(bookDesc, prometheus.GaugeValue, float64(res.BookUnparsedCount), "unparsed")

	metr <- prometheus.MustNewConstMetric(pageDesc, prometheus.GaugeValue, float64(res.PageCount), "all")
	metr <- prometheus.MustNewConstMetric(pageDesc, prometheus.GaugeValue, float64(res.PageUnloadedCount), "unloaded")
	metr <- prometheus.MustNewConstMetric(pageDesc, prometheus.GaugeValue, float64(res.PageWithoutBodyCount), "no_body")

	metr <- prometheus.MustNewConstMetric(fileDesc, prometheus.GaugeValue, float64(res.PageFileSize), "page")
	metr <- prometheus.MustNewConstMetric(fileDesc, prometheus.GaugeValue, float64(res.FileSize), "fs")

	for _, worker := range res.Workers {
		metr <- prometheus.MustNewConstMetric(workerDesc, prometheus.GaugeValue, float64(worker.InQueueCount), worker.Name, "in_queue")
		metr <- prometheus.MustNewConstMetric(workerDesc, prometheus.GaugeValue, float64(worker.InWorkCount), worker.Name, "in_work")
		metr <- prometheus.MustNewConstMetric(workerDesc, prometheus.GaugeValue, float64(worker.RunnersCount), worker.Name, "runners")
	}
}

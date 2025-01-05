package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"hgnext/internal/entities"
)

type infoProvider interface {
	SystemInfo(ctx context.Context) (entities.SystemSizeInfoWithMonitor, error)
}

var (
	bookTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: SystemName,
		Subsystem: SubSystemName,
		Name:      "book_total",
		Help:      "Количество книг по статусам",
	}, []string{"type"})
	pageTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: SystemName,
		Subsystem: SubSystemName,
		Name:      "page_total",
		Help:      "Количество страниц по статусам",
	}, []string{"type"})
	fileTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: SystemName,
		Subsystem: SubSystemName,
		Name:      "file_total",
		Help:      "Количество файлов по статусам",
	}, []string{"type"})
	fileBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: SystemName,
		Subsystem: SubSystemName,
		Name:      "file_bytes",
		Help:      "Размер файлов по статусам",
	}, []string{"type"})
	workerTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: SystemName,
		Subsystem: SubSystemName,
		Name:      "worker_total",
		Help:      "Данные воркеров",
	}, []string{"worker_name", "counter"})
	lastCollectorScrapeDuration = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: SystemName,
		Subsystem: SubSystemName,
		Name:      "info_scrape_duration_seconds",
		Help:      "Время последней сборки основных данных",
	})
)

type SystemInfoCollector struct {
	logger       *slog.Logger
	infoProvider infoProvider
	interval     time.Duration
}

func NewSystemInfoCollector(
	logger *slog.Logger,
	infoProvider infoProvider,
	interval time.Duration,
) *SystemInfoCollector {
	return &SystemInfoCollector{
		logger:       logger,
		infoProvider: infoProvider,
		interval:     interval,
	}
}

func (c *SystemInfoCollector) Name() string {
	return "system info collector"
}

func (c *SystemInfoCollector) Start(ctx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		c.logger.InfoContext(ctx, "system info collector start")
		defer c.logger.InfoContext(ctx, "system info collector stop")

		ticker := time.NewTicker(c.interval)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				c.collect(ctx)
			}
		}
	}()

	return done, nil
}

//nolint:promlinter // ложно-положительное срабатывание
func (c *SystemInfoCollector) collect(ctx context.Context) {
	tStart := time.Now()

	res, err := c.infoProvider.SystemInfo(ctx)
	if err != nil {
		c.logger.ErrorContext(
			ctx, "failed scrap system info",
			slog.Any("error", err),
		)

		return
	}

	lastCollectorScrapeDuration.Set(time.Since(tStart).Seconds())

	bookTotal.WithLabelValues("all").Set(float64(res.BookCount))
	bookTotal.WithLabelValues("downloaded").Set(float64(res.DownloadedBookCount))
	bookTotal.WithLabelValues("verified").Set(float64(res.VerifiedBookCount))
	bookTotal.WithLabelValues("rebuilded").Set(float64(res.RebuildedBookCount))
	bookTotal.WithLabelValues("unparsed").Set(float64(res.BookUnparsedCount))
	bookTotal.WithLabelValues("deleted").Set(float64(res.DeletedBookCount))

	pageTotal.WithLabelValues("all").Set(float64(res.PageCount))
	pageTotal.WithLabelValues("unloaded").Set(float64(res.PageUnloadedCount))
	pageTotal.WithLabelValues("no_body").Set(float64(res.PageWithoutBodyCount))
	pageTotal.WithLabelValues("deleted").Set(float64(res.DeletedPageCount))

	fileTotal.WithLabelValues("all").Set(float64(res.FileCount))
	fileTotal.WithLabelValues("unhashed").Set(float64(res.UnhashedFileCount))
	fileTotal.WithLabelValues("dead_hash").Set(float64(res.DeadHashCount))

	fileBytes.WithLabelValues("page").Set(float64(res.PageFileSize))
	fileBytes.WithLabelValues("fs").Set(float64(res.FileSize))

	for _, worker := range res.Workers {
		workerTotal.WithLabelValues(worker.Name, "in_queue").Set(float64(worker.InQueueCount))
		workerTotal.WithLabelValues(worker.Name, "in_work").Set(float64(worker.InWorkCount))
		workerTotal.WithLabelValues(worker.Name, "runners").Set(float64(worker.RunnersCount))
	}
}

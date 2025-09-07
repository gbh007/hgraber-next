package metrics

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

type infoProvider interface {
	WorkersInfo(ctx context.Context) []systemmodel.SystemWorkerStat
}

type statisticProvider interface {
	BooksCountByAuthor(ctx context.Context) (map[string]int64, error)
	PageSizeByAuthor(ctx context.Context) (map[string]core.SizeWithCount, error)
	BookSizes(ctx context.Context) (map[uuid.UUID]core.SizeWithCount, error)
	SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error)
}

type config interface {
	MainInfo() time.Duration
	BookStatistic() time.Duration
}

var (
	bookTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricserver.BookTotalName,
		Help: "Количество книг по статусам",
	}, []string{metricserver.TypeLabel})
	pageTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricserver.PageTotalName,
		Help: "Количество страниц по статусам",
	}, []string{metricserver.TypeLabel})
	fileTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricserver.FileTotalName,
		Help: "Количество файлов по статусам",
	}, []string{metricserver.TypeLabel, metriccore.FSIDLabel})
	fileBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricserver.FileBytesName,
		Help: "Размер файлов по статусам",
	}, []string{metricserver.TypeLabel, metriccore.FSIDLabel})
	lastCollectorScrapeDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricserver.LastCollectorScrapeDurationName,
		Help: "Время последней сборки данных",
	}, []string{metricserver.ScrapeNameLabel})
)

// TODO: это по факту уже контроллер, надо вынести в контроллеры
type SystemInfoCollector struct {
	logger            *slog.Logger
	infoProvider      infoProvider
	statisticProvider statisticProvider
	mainInfoInterval  time.Duration
	statisticInterval time.Duration

	statisticCollector *BookStatisticCollector
}

func NewSystemInfoCollector(
	logger *slog.Logger,
	infoProvider infoProvider,
	statisticProvider statisticProvider,
	config config,
) (*SystemInfoCollector, error) {
	bookStatisticCollector := NewBookStatisticCollector()

	if config.BookStatistic() > 0 {
		err := prometheus.Register(bookStatisticCollector)
		if err != nil {
			return nil, fmt.Errorf("register book collector: %w", err)
		}
	}

	err := RegisterWorkerInfoCollector(
		logger,
		infoProvider,
		time.Millisecond*100, // TODO: настраивать таймаут через конфиг
	)
	if err != nil {
		return nil, fmt.Errorf("register worker collector: %w", err)
	}

	return &SystemInfoCollector{
		logger:             logger,
		infoProvider:       infoProvider,
		statisticProvider:  statisticProvider,
		mainInfoInterval:   config.MainInfo(),
		statisticInterval:  config.BookStatistic(),
		statisticCollector: bookStatisticCollector,
	}, nil
}

func (c *SystemInfoCollector) Name() string {
	return "system info collector"
}

func (c *SystemInfoCollector) Start(ctx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	wg := new(sync.WaitGroup)

	if c.mainInfoInterval > 0 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			c.logger.InfoContext(ctx, "system info collector start")
			defer c.logger.InfoContext(ctx, "system info collector stop")

			c.collectMainInfo(ctx)

			ticker := time.NewTicker(c.mainInfoInterval)

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					c.collectMainInfo(ctx)
				}
			}
		}()
	}

	if c.statisticInterval > 0 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			c.logger.InfoContext(ctx, "book statistic collector start")
			defer c.logger.InfoContext(ctx, "book statistic collector stop")

			c.collectBookSizeStatistic(ctx)
			c.collectBookCountByAuthorStatistic(ctx)
			c.collectPageSizeByAuthorStatistic(ctx)

			ticker := time.NewTicker(c.statisticInterval)

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					c.collectBookSizeStatistic(ctx)
					c.collectBookCountByAuthorStatistic(ctx)
					c.collectPageSizeByAuthorStatistic(ctx)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	return done, nil
}

func (c *SystemInfoCollector) collectMainInfo(ctx context.Context) {
	tStart := time.Now()

	res, err := c.statisticProvider.SystemSize(ctx)
	if err != nil {
		c.logger.ErrorContext(
			ctx, "failed scrap system info",
			slog.Any("error", err),
		)

		return
	}

	lastCollectorScrapeDuration.WithLabelValues("main_info").Set(time.Since(tStart).Seconds())

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

	fileTotal.WithLabelValues("dead_hash", "").Set(float64(res.DeadHashCount))

	for fsID, v := range res.FileCountByFS {
		fileTotal.WithLabelValues("all", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.UnhashedFileCountByFS {
		fileTotal.WithLabelValues("unhashed", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.InvalidFileCountByFS {
		fileTotal.WithLabelValues("invalid", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.DetachedFileCountByFS {
		fileTotal.WithLabelValues("detached", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.PageFileSizeByFS {
		fileBytes.WithLabelValues("page", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.FileSizeByFS {
		fileBytes.WithLabelValues("fs", fsID.String()).Set(float64(v))
	}
}

func (c *SystemInfoCollector) collectBookSizeStatistic(ctx context.Context) {
	tStart := time.Now()

	bookSizes, err := c.statisticProvider.BookSizes(ctx)
	if err != nil {
		c.logger.ErrorContext(
			ctx, "failed scrap book size statistic",
			slog.Any("error", err),
		)

		return
	}

	lastCollectorScrapeDuration.WithLabelValues("book_size").Set(time.Since(tStart).Seconds())

	pageInBookBuckets := make(map[float64]uint64, len(pageInBookBucket))
	bookSizeBuckets := make(map[float64]uint64, len(bookSizeBucket))
	pageSizeBuckets := make(map[float64]uint64, len(pageSizeBucket))

	var (
		pageInBookCount uint64
		pageInBookSum   float64

		bookSizeCount uint64
		bookSizeSum   float64

		pageSizeCount uint64
		pageSizeSum   float64
	)

	for _, size := range bookSizes {
		pageInBookCount++
		pageInBookSum += float64(size.Count)

		bookSizeCount++
		bookSizeSum += float64(size.Size)

		pageAvgSize := float64(size.Size) / float64(size.Count)

		pageSizeCount += uint64(size.Count)
		pageSizeSum += pageAvgSize

		for _, bucket := range pageInBookBucket {
			if size.Count <= int64(bucket) {
				pageInBookBuckets[bucket]++
			}
		}

		for _, bucket := range bookSizeBucket {
			if size.Size <= int64(bucket) {
				bookSizeBuckets[bucket]++
			}
		}

		for _, bucket := range pageSizeBucket {
			if pageAvgSize <= bucket {
				pageSizeBuckets[bucket] += uint64(size.Count)
			}
		}
	}

	c.statisticCollector.bookPages.Store(&histogramData{
		buckets: pageInBookBuckets,
		count:   pageInBookCount,
		sum:     pageInBookSum,
	})

	c.statisticCollector.bookSize.Store(&histogramData{
		buckets: bookSizeBuckets,
		count:   bookSizeCount,
		sum:     bookSizeSum,
	})

	c.statisticCollector.pageSize.Store(&histogramData{
		buckets: pageSizeBuckets,
		count:   pageSizeCount,
		sum:     pageSizeSum,
	})
}

func (c *SystemInfoCollector) collectBookCountByAuthorStatistic(ctx context.Context) {
	tStart := time.Now()

	bookCountByAuthor, err := c.statisticProvider.BooksCountByAuthor(ctx)
	if err != nil {
		c.logger.ErrorContext(
			ctx, "failed scrap book by author statistic",
			slog.Any("error", err),
		)

		return
	}

	lastCollectorScrapeDuration.WithLabelValues("book_by_author").Set(time.Since(tStart).Seconds())

	bookCountByAuthorBuckets := make(map[float64]uint64, len(bookCountByAuthorBucket))

	var (
		bookCountByAuthorCount uint64
		bookCountByAuthorSum   float64
	)

	for _, size := range bookCountByAuthor {
		bookCountByAuthorCount++
		bookCountByAuthorSum += float64(size)

		for _, bucket := range bookCountByAuthorBucket {
			if size <= int64(bucket) {
				bookCountByAuthorBuckets[bucket]++
			}
		}
	}

	c.statisticCollector.bookCountByAuthor.Store(&histogramData{
		buckets: bookCountByAuthorBuckets,
		count:   bookCountByAuthorCount,
		sum:     bookCountByAuthorSum,
	})
}

func (c *SystemInfoCollector) collectPageSizeByAuthorStatistic(ctx context.Context) {
	tStart := time.Now()

	pagesSizeByAuthors, err := c.statisticProvider.PageSizeByAuthor(ctx)
	if err != nil {
		c.logger.ErrorContext(
			ctx, "failed scrap page size by author statistic",
			slog.Any("error", err),
		)

		return
	}

	lastCollectorScrapeDuration.WithLabelValues("page_size_bu_author").Set(time.Since(tStart).Seconds())

	pagesByAuthorBuckets := make(map[float64]uint64, len(pagesByAuthorBucket))
	pagesSizeByAuthorBuckets := make(map[float64]uint64, len(pagesSizeByAuthorBucket))

	var (
		pagesByAuthorCount uint64
		pagesByAuthorSum   float64

		pagesSizeByAuthorCount uint64
		pagesSizeByAuthorSum   float64
	)

	for _, size := range pagesSizeByAuthors {
		pagesByAuthorCount++
		pagesByAuthorSum += float64(size.Count)

		pagesSizeByAuthorCount++
		pagesSizeByAuthorSum += float64(size.Size)

		for _, bucket := range pagesByAuthorBucket {
			if size.Count <= int64(bucket) {
				pagesByAuthorBuckets[bucket]++
			}
		}

		for _, bucket := range pagesSizeByAuthorBucket {
			if size.Size <= int64(bucket) {
				pagesSizeByAuthorBuckets[bucket]++
			}
		}
	}

	c.statisticCollector.pagesByAuthor.Store(&histogramData{
		buckets: pagesByAuthorBuckets,
		count:   pagesByAuthorCount,
		sum:     pagesByAuthorSum,
	})

	c.statisticCollector.pagesSizeByAuthor.Store(&histogramData{
		buckets: pagesSizeByAuthorBuckets,
		count:   pagesSizeByAuthorCount,
		sum:     pagesSizeByAuthorSum,
	})
}

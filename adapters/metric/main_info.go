package metric

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
	"github.com/gbh007/hgraber-next/adapters/metric/metricstatistic"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
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

// TODO: это по факту уже контроллер, надо вынести в контроллеры
type SystemInfoCollector struct {
	logger            *slog.Logger
	infoProvider      infoProvider
	statisticProvider statisticProvider
	mainInfoInterval  time.Duration
	statisticInterval time.Duration

	statisticCollector *metricstatistic.BookStatisticCollector
}

func (mp MetricProvider) NewSystemInfoCollector(
	logger *slog.Logger,
	infoProvider infoProvider,
	statisticProvider statisticProvider,
	config config,
) (*SystemInfoCollector, error) {
	bookStatisticCollector := metricstatistic.NewBookStatisticCollector()

	if config.BookStatistic() > 0 {
		err := mp.registerer.Register(bookStatisticCollector)
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

	metricserver.LastCollectorScrapeDuration.WithLabelValues("main_info").Set(time.Since(tStart).Seconds())

	metricserver.BookTotal.WithLabelValues("all").Set(float64(res.BookCount))
	metricserver.BookTotal.WithLabelValues("downloaded").Set(float64(res.DownloadedBookCount))
	metricserver.BookTotal.WithLabelValues("verified").Set(float64(res.VerifiedBookCount))
	metricserver.BookTotal.WithLabelValues("rebuilded").Set(float64(res.RebuildedBookCount))
	metricserver.BookTotal.WithLabelValues("unparsed").Set(float64(res.BookUnparsedCount))
	metricserver.BookTotal.WithLabelValues("deleted").Set(float64(res.DeletedBookCount))

	metricserver.PageTotal.WithLabelValues("all").Set(float64(res.PageCount))
	metricserver.PageTotal.WithLabelValues("unloaded").Set(float64(res.PageUnloadedCount))
	metricserver.PageTotal.WithLabelValues("no_body").Set(float64(res.PageWithoutBodyCount))
	metricserver.PageTotal.WithLabelValues("deleted").Set(float64(res.DeletedPageCount))

	metricfs.FileTotal.WithLabelValues("dead_hash", "").Set(float64(res.DeadHashCount))

	for fsID, v := range res.FileCountByFS {
		metricfs.FileTotal.WithLabelValues("all", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.UnhashedFileCountByFS {
		metricfs.FileTotal.WithLabelValues("unhashed", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.InvalidFileCountByFS {
		metricfs.FileTotal.WithLabelValues("invalid", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.DetachedFileCountByFS {
		metricfs.FileTotal.WithLabelValues("detached", fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.PageFileSizeByFS {
		metricfs.FileBytes.WithLabelValues(metricfs.TypeLabelValuePage, fsID.String()).Set(float64(v))
	}

	for fsID, v := range res.FileSizeByFS {
		metricfs.FileBytes.WithLabelValues(metricfs.TypeLabelValueFS, fsID.String()).Set(float64(v))
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

	metricserver.LastCollectorScrapeDuration.WithLabelValues("book_size").Set(time.Since(tStart).Seconds())

	pageInBookBuckets := make(map[float64]uint64, len(metricstatistic.PageInBookBucket))
	bookSizeBuckets := make(map[float64]uint64, len(metricstatistic.BookSizeBucket))
	pageSizeBuckets := make(map[float64]uint64, len(metricstatistic.PageSizeBucket))

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

		for _, bucket := range metricstatistic.PageInBookBucket {
			if size.Count <= int64(bucket) {
				pageInBookBuckets[bucket]++
			}
		}

		for _, bucket := range metricstatistic.BookSizeBucket {
			if size.Size <= int64(bucket) {
				bookSizeBuckets[bucket]++
			}
		}

		for _, bucket := range metricstatistic.PageSizeBucket {
			if pageAvgSize <= bucket {
				pageSizeBuckets[bucket] += uint64(size.Count)
			}
		}
	}

	c.statisticCollector.BookPages.Store(&metricstatistic.HistogramData{
		Buckets: pageInBookBuckets,
		Count:   pageInBookCount,
		Sum:     pageInBookSum,
	})

	c.statisticCollector.BookSize.Store(&metricstatistic.HistogramData{
		Buckets: bookSizeBuckets,
		Count:   bookSizeCount,
		Sum:     bookSizeSum,
	})

	c.statisticCollector.PageSize.Store(&metricstatistic.HistogramData{
		Buckets: pageSizeBuckets,
		Count:   pageSizeCount,
		Sum:     pageSizeSum,
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

	metricserver.LastCollectorScrapeDuration.WithLabelValues("book_by_author").Set(time.Since(tStart).Seconds())

	bookCountByAuthorBuckets := make(map[float64]uint64, len(metricstatistic.BookCountByAuthorBucket))

	var (
		bookCountByAuthorCount uint64
		bookCountByAuthorSum   float64
	)

	for _, size := range bookCountByAuthor {
		bookCountByAuthorCount++
		bookCountByAuthorSum += float64(size)

		for _, bucket := range metricstatistic.BookCountByAuthorBucket {
			if size <= int64(bucket) {
				bookCountByAuthorBuckets[bucket]++
			}
		}
	}

	c.statisticCollector.BookCountByAuthor.Store(&metricstatistic.HistogramData{
		Buckets: bookCountByAuthorBuckets,
		Count:   bookCountByAuthorCount,
		Sum:     bookCountByAuthorSum,
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

	metricserver.LastCollectorScrapeDuration.WithLabelValues("page_size_by_author").Set(time.Since(tStart).Seconds())

	pagesByAuthorBuckets := make(map[float64]uint64, len(metricstatistic.PagesByAuthorBucket))
	pagesSizeByAuthorBuckets := make(map[float64]uint64, len(metricstatistic.PagesSizeByAuthorBucket))

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

		for _, bucket := range metricstatistic.PagesByAuthorBucket {
			if size.Count <= int64(bucket) {
				pagesByAuthorBuckets[bucket]++
			}
		}

		for _, bucket := range metricstatistic.PagesSizeByAuthorBucket {
			if size.Size <= int64(bucket) {
				pagesSizeByAuthorBuckets[bucket]++
			}
		}
	}

	c.statisticCollector.PagesByAuthor.Store(&metricstatistic.HistogramData{
		Buckets: pagesByAuthorBuckets,
		Count:   pagesByAuthorCount,
		Sum:     pagesByAuthorSum,
	})

	c.statisticCollector.PagesSizeByAuthor.Store(&metricstatistic.HistogramData{
		Buckets: pagesSizeByAuthorBuckets,
		Count:   pagesSizeByAuthorCount,
		Sum:     pagesSizeByAuthorSum,
	})
}

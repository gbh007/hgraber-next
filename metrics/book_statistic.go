package metrics

import (
	"sync/atomic"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	kilobyte = 1 << 10
	megabyte = 1 << 20
	gigabyte = 1 << 30
)

var (
	pageInBookDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metriccore.SystemName, metriccore.SubSystemName, "statistic_page_in_book"),
		"Данные количества страниц в книге",
		nil, nil,
	)
	pageInBookBucket = []float64{5, 10, 20, 30, 50, 100, 250, 500, 1000, 2000}
	bookSizeDesc     = prometheus.NewDesc(
		prometheus.BuildFQName(metriccore.SystemName, metriccore.SubSystemName, "statistic_book_size"),
		"Данные размера книг",
		nil, nil,
	)
	bookSizeBucket = []float64{
		megabyte,
		5 * megabyte,
		10 * megabyte,
		30 * megabyte,
		50 * megabyte,
		100 * megabyte,
		300 * megabyte,
		600 * megabyte,
		gigabyte,
	}
	pageSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metriccore.SystemName, metriccore.SubSystemName, "statistic_page_size"),
		"Данные размера страниц",
		nil, nil,
	)
	pageSizeBucket = []float64{
		50 * kilobyte,
		200 * kilobyte,
		500 * kilobyte,
		megabyte,
		2 * megabyte,
		5 * megabyte,
		10 * megabyte,
	}
	bookCountByAuthorDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metriccore.SystemName, metriccore.SubSystemName, "statistic_books_by_author"),
		"Данные количества книг у одного автора",
		nil, nil,
	)
	bookCountByAuthorBucket = []float64{5, 10, 30, 50, 100, 250, 500}
	pagesByAuthorDesc       = prometheus.NewDesc(
		prometheus.BuildFQName(metriccore.SystemName, metriccore.SubSystemName, "statistic_pages_by_author"),
		"Данные количества страниц у одного автора",
		nil, nil,
	)
	pagesByAuthorBucket   = []float64{50, 100, 500, 1000, 5_000, 10_000, 20_000}
	pagesSizeByAuthorDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metriccore.SystemName, metriccore.SubSystemName, "statistic_pages_size_by_author"),
		"Данные размера страниц у одного автора",
		nil, nil,
	)
	pagesSizeByAuthorBucket = []float64{
		50 * megabyte,
		200 * megabyte,
		500 * megabyte,
		gigabyte,
		2 * gigabyte,
		5 * gigabyte,
		10 * gigabyte,
	}
)

var _ prometheus.Collector = (*BookStatisticCollector)(nil)

type histogramData struct {
	buckets map[float64]uint64
	count   uint64
	sum     float64
}

type BookStatisticCollector struct {
	bookPages         atomic.Pointer[histogramData]
	bookSize          atomic.Pointer[histogramData]
	pageSize          atomic.Pointer[histogramData]
	bookCountByAuthor atomic.Pointer[histogramData]
	pagesByAuthor     atomic.Pointer[histogramData]
	pagesSizeByAuthor atomic.Pointer[histogramData]
}

func NewBookStatisticCollector() *BookStatisticCollector {
	return &BookStatisticCollector{
		bookPages:         atomic.Pointer[histogramData]{},
		bookSize:          atomic.Pointer[histogramData]{},
		pageSize:          atomic.Pointer[histogramData]{},
		bookCountByAuthor: atomic.Pointer[histogramData]{},
		pagesByAuthor:     atomic.Pointer[histogramData]{},
		pagesSizeByAuthor: atomic.Pointer[histogramData]{},
	}
}

func (c *BookStatisticCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- pageInBookDesc

	desc <- bookSizeDesc

	desc <- pageSizeDesc

	desc <- bookCountByAuthorDesc

	desc <- pagesByAuthorDesc

	desc <- pagesSizeByAuthorDesc
}

func (c *BookStatisticCollector) Collect(metr chan<- prometheus.Metric) {
	bookPages := c.bookPages.Load()
	if bookPages != nil {
		metr <- prometheus.MustNewConstHistogram(
			pageInBookDesc,
			bookPages.count,
			bookPages.sum,
			bookPages.buckets,
		)
	}

	bookSize := c.bookSize.Load()
	if bookSize != nil {
		metr <- prometheus.MustNewConstHistogram(
			bookSizeDesc,
			bookSize.count,
			bookSize.sum,
			bookSize.buckets,
		)
	}

	pageSize := c.pageSize.Load()
	if pageSize != nil {
		metr <- prometheus.MustNewConstHistogram(
			pageSizeDesc,
			pageSize.count,
			pageSize.sum,
			pageSize.buckets,
		)
	}

	bookCountByAuthor := c.bookCountByAuthor.Load()
	if bookCountByAuthor != nil {
		metr <- prometheus.MustNewConstHistogram(
			bookCountByAuthorDesc,
			bookCountByAuthor.count,
			bookCountByAuthor.sum,
			bookCountByAuthor.buckets,
		)
	}

	pagesByAuthor := c.pagesByAuthor.Load()
	if pagesByAuthor != nil {
		metr <- prometheus.MustNewConstHistogram(
			pagesByAuthorDesc,
			pagesByAuthor.count,
			pagesByAuthor.sum,
			pagesByAuthor.buckets,
		)
	}

	pagesSizeByAuthor := c.pagesSizeByAuthor.Load()
	if pagesSizeByAuthor != nil {
		metr <- prometheus.MustNewConstHistogram(
			pagesSizeByAuthorDesc,
			pagesSizeByAuthor.count,
			pagesSizeByAuthor.sum,
			pagesSizeByAuthor.buckets,
		)
	}
}

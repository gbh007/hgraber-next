package metricstatistic

import (
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

var (
	pageInBookDesc = prometheus.NewDesc(
		PageInBook,
		"Данные количества страниц в книге",
		nil, nil,
	)
	PageInBookBucket = []float64{5, 10, 20, 30, 50, 100, 250, 500, 1000, 2000}
	bookSizeDesc     = prometheus.NewDesc(
		BookSize,
		"Данные размера книг",
		nil, nil,
	)
	BookSizeBucket = []float64{
		metriccore.Megabyte,
		5 * metriccore.Megabyte,
		10 * metriccore.Megabyte,
		30 * metriccore.Megabyte,
		50 * metriccore.Megabyte,
		100 * metriccore.Megabyte,
		300 * metriccore.Megabyte,
		600 * metriccore.Megabyte,
		metriccore.Gigabyte,
	}
	pageSizeDesc = prometheus.NewDesc(
		PageSize,
		"Данные размера страниц",
		nil, nil,
	)
	PageSizeBucket = []float64{
		50 * metriccore.Kilobyte,
		200 * metriccore.Kilobyte,
		500 * metriccore.Kilobyte,
		metriccore.Megabyte,
		2 * metriccore.Megabyte,
		5 * metriccore.Megabyte,
		10 * metriccore.Megabyte,
	}
	bookCountByAuthorDesc = prometheus.NewDesc(
		BookCountByAuthor,
		"Данные количества книг у одного автора",
		nil, nil,
	)
	BookCountByAuthorBucket = []float64{5, 10, 30, 50, 100, 250, 500}
	pagesByAuthorDesc       = prometheus.NewDesc(
		PagesByAuthor,
		"Данные количества страниц у одного автора",
		nil, nil,
	)
	PagesByAuthorBucket   = []float64{50, 100, 500, 1000, 5_000, 10_000, 20_000}
	pagesSizeByAuthorDesc = prometheus.NewDesc(
		PagesSizeByAuthor,
		"Данные размера страниц у одного автора",
		nil, nil,
	)
	PagesSizeByAuthorBucket = []float64{
		50 * metriccore.Megabyte,
		200 * metriccore.Megabyte,
		500 * metriccore.Megabyte,
		metriccore.Gigabyte,
		2 * metriccore.Gigabyte,
		5 * metriccore.Gigabyte,
		10 * metriccore.Gigabyte,
	}
)

var _ prometheus.Collector = (*BookStatisticCollector)(nil)

type HistogramData struct {
	Buckets map[float64]uint64
	Count   uint64
	Sum     float64
}

type BookStatisticCollector struct {
	BookPages         atomic.Pointer[HistogramData]
	BookSize          atomic.Pointer[HistogramData]
	PageSize          atomic.Pointer[HistogramData]
	BookCountByAuthor atomic.Pointer[HistogramData]
	PagesByAuthor     atomic.Pointer[HistogramData]
	PagesSizeByAuthor atomic.Pointer[HistogramData]
}

func NewBookStatisticCollector() *BookStatisticCollector {
	return &BookStatisticCollector{
		BookPages:         atomic.Pointer[HistogramData]{},
		BookSize:          atomic.Pointer[HistogramData]{},
		PageSize:          atomic.Pointer[HistogramData]{},
		BookCountByAuthor: atomic.Pointer[HistogramData]{},
		PagesByAuthor:     atomic.Pointer[HistogramData]{},
		PagesSizeByAuthor: atomic.Pointer[HistogramData]{},
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
	bookPages := c.BookPages.Load()
	if bookPages != nil {
		metr <- prometheus.MustNewConstHistogram(
			pageInBookDesc,
			bookPages.Count,
			bookPages.Sum,
			bookPages.Buckets,
		)
	}

	bookSize := c.BookSize.Load()
	if bookSize != nil {
		metr <- prometheus.MustNewConstHistogram(
			bookSizeDesc,
			bookSize.Count,
			bookSize.Sum,
			bookSize.Buckets,
		)
	}

	pageSize := c.PageSize.Load()
	if pageSize != nil {
		metr <- prometheus.MustNewConstHistogram(
			pageSizeDesc,
			pageSize.Count,
			pageSize.Sum,
			pageSize.Buckets,
		)
	}

	bookCountByAuthor := c.BookCountByAuthor.Load()
	if bookCountByAuthor != nil {
		metr <- prometheus.MustNewConstHistogram(
			bookCountByAuthorDesc,
			bookCountByAuthor.Count,
			bookCountByAuthor.Sum,
			bookCountByAuthor.Buckets,
		)
	}

	pagesByAuthor := c.PagesByAuthor.Load()
	if pagesByAuthor != nil {
		metr <- prometheus.MustNewConstHistogram(
			pagesByAuthorDesc,
			pagesByAuthor.Count,
			pagesByAuthor.Sum,
			pagesByAuthor.Buckets,
		)
	}

	pagesSizeByAuthor := c.PagesSizeByAuthor.Load()
	if pagesSizeByAuthor != nil {
		metr <- prometheus.MustNewConstHistogram(
			pagesSizeByAuthorDesc,
			pagesSizeByAuthor.Count,
			pagesSizeByAuthor.Sum,
			pagesSizeByAuthor.Buckets,
		)
	}
}

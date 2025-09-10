package metricstatistic

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
)

var (
	PageInBook = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"statistic_page_in_book",
	)
	BookSize = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"statistic_book_size",
	)
	PageSize = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"statistic_page_size",
	)
	BookCountByAuthor = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"statistic_books_by_author",
	)
	PagesByAuthor = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"statistic_pages_by_author",
	)
	PagesSizeByAuthor = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"statistic_pages_size_by_author",
	)
)

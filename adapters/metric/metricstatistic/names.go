package metricstatistic

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

const SubSystemName = "statistic"

var (
	PageInBook = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"page_in_book",
	)
	BookSize = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"book_size",
	)
	PageSize = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"page_size",
	)
	BookCountByAuthor = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"books_by_author",
	)
	PagesByAuthor = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"pages_by_author",
	)
	PagesSizeByAuthor = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"pages_size_by_author",
	)
)

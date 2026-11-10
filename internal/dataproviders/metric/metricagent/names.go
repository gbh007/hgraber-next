package metricagent

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
)

const (
	SubSystemName   = "agent"
	ParserNameLabel = "parser_name"
)

var (
	ParserActionSecondsName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"parser_action_seconds",
	)
	WebCacheTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"web_cache_total",
	)
	WebCacheBytesName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"web_cache_bytes",
	)
	WebCacheFileTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"web_cache_file_total",
	)
)

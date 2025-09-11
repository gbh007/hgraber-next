package metricagent

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
)

const (
	SubSystemName   = "agent"
	ParserNameLabel = "parser_name"
)

var (
	FSActionSecondsName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"fs_action_seconds",
	)
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
)

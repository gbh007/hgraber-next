package metricfs

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

const SubSystemName = "fs"

var (
	FileTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"file_total",
	)
	FileBytesName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"file_bytes",
	)
	ActionSecondsName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"action_seconds",
	)
)

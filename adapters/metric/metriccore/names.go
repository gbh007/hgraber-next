package metriccore

import "github.com/prometheus/client_golang/prometheus"

var VersionInfoName = prometheus.BuildFQName(
	SystemName,
	"",
	"version_info",
)

package metricdatabase

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

const (
	SubSystemName = "database"
)

var (
	ActiveRequestName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"active_request",
	)
	OpenConnectionName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"open_connection",
	)
	RequestDurationName = prometheus.BuildFQName(
		metriccore.SystemName,
		SubSystemName,
		"request_duration",
	)
)

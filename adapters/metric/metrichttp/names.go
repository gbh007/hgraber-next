package metrichttp

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

const (
	ServerSystemName    = "http_server"
	ServerAddrLabelName = "server_addr"
	OperationLabelName  = "operation"
	StatusLabelName     = "status"
)

var (
	ServerHandleDurationName = prometheus.BuildFQName(
		metriccore.SystemName,
		ServerSystemName,
		"handle_duration",
	)
	ServerActiveHandlersTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		ServerSystemName,
		"active_handlers_total",
	)
)

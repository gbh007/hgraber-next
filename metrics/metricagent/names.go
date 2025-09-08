package metricagent

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
)

const (
	SubSystemName = "agent"
)

var FSActionSecondsName = prometheus.BuildFQName(
	metriccore.SystemName,
	SubSystemName,
	"fs_action_seconds",
)

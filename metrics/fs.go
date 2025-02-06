package metrics

import (
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var fsActionTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace: SystemName,
	Subsystem: SubSystemName,
	Name:      "fs_action_seconds",
	Help:      "Время действий с файловой системой",
}, []string{"action", "fs_id"})

func (MetricProvider) RegisterFSActionTime(action string, fsID *uuid.UUID, d time.Duration) {
	var fs string

	if fsID != nil {
		fs = fsID.String()
	}

	fsActionTime.WithLabelValues(action, fs).Observe(d.Seconds())
}

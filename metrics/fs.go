package metrics

import (
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

var fsActionTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: metricserver.FSActionSecondsName,
	Help: "Время действий с файловой системой",
}, []string{metriccore.ActionLabel, metriccore.FSIDLabel})

func (MetricProvider) RegisterFSActionTime(action string, fsID *uuid.UUID, d time.Duration) {
	var fs string

	if fsID != nil {
		fs = fsID.String()
	}

	fsActionTime.WithLabelValues(action, fs).Observe(d.Seconds())
}

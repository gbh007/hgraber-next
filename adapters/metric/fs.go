package metric

import (
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func (MetricProvider) RegisterFSActionTime(action string, fsID *uuid.UUID, d time.Duration) {
	var fs string

	if fsID != nil {
		fs = fsID.String()
	}

	metricfs.ActionTime.WithLabelValues(action, fs).Observe(d.Seconds())
}

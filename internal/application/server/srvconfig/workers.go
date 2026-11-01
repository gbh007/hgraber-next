package srvconfig

import (
	"time"
)

type ServerWorker struct {
	Name      string        `toml:"name"`
	Count     int32         `toml:"count"`
	QueueSize int           `toml:"queue_size"`
	Interval  time.Duration `toml:"interval"`
}

func (w ServerWorker) GetName() string {
	return w.Name
}

func (w ServerWorker) GetCount() int32 {
	return w.Count
}

func (w ServerWorker) GetQueueSize() int {
	return w.QueueSize
}

func (w ServerWorker) GetInterval() time.Duration {
	return w.Interval
}

package config

import (
	"time"
)

type ServerWorker struct {
	Name      string        `toml:"name" yaml:"name" envconfig:"NAME"`
	Count     int32         `toml:"count" yaml:"count" envconfig:"COUNT"`
	QueueSize int           `toml:"queue_size" yaml:"queue_size" envconfig:"QUEUE_SIZE"`
	Interval  time.Duration `toml:"interval" yaml:"interval" envconfig:"INTERVAL"`
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

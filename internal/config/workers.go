package config

import "time"

type Workers struct {
	Page          Worker `yaml:"page" envconfig:"PAGE"`
	Book          Worker `yaml:"book" envconfig:"BOOK"`
	Hasher        Worker `yaml:"hasher" envconfig:"HASHER"`
	Exporter      Worker `yaml:"exporter" envconfig:"EXPORTER"`
	Tasker        Worker `yaml:"tasker" envconfig:"TASKER"`
	FileValidator Worker `yaml:"file_validator" envconfig:"FILE_VALIDATOR"`
}

func WorkersDefault() Workers {
	return Workers{
		Page:          WorkerDefault(),
		Book:          WorkerDefault(),
		Hasher:        WorkerDefault(),
		Exporter:      WorkerDefault(),
		Tasker:        WorkerDefault(),
		FileValidator: WorkerDefault(),
	}
}

type Worker struct {
	Count     int32         `yaml:"count" envconfig:"COUNT"`
	QueueSize int           `yaml:"queue_size" envconfig:"QUEUE_SIZE"`
	Interval  time.Duration `yaml:"interval" envconfig:"INTERVAL"`
}

func WorkerDefault() Worker {
	return Worker{
		Count:     1,
		QueueSize: 100,
		Interval:  time.Minute,
	}
}

func (w Worker) GetCount() int32 {
	return w.Count
}

func (w Worker) GetQueueSize() int {
	return w.QueueSize
}

func (w Worker) GetInterval() time.Duration {
	return w.Interval
}

package config

import "time"

type Workers struct {
	Page                   Worker `yaml:"page" envconfig:"PAGE"`
	Book                   Worker `yaml:"book" envconfig:"BOOK"`
	Hasher                 Worker `yaml:"hasher" envconfig:"HASHER"`
	Exporter               Worker `yaml:"exporter" envconfig:"EXPORTER"`
	Tasker                 Worker `yaml:"tasker" envconfig:"TASKER"`
	FileValidator          Worker `yaml:"file_validator" envconfig:"FILE_VALIDATOR"`
	FileTransferer         Worker `yaml:"file_transferer" envconfig:"FILE_TRANSFERER"`
	MassloadSizer          Worker `yaml:"massload_sizer" envconfig:"MASSLOAD_SIZER"`
	MassloadAttributeSizer Worker `yaml:"massload_attribute_sizer" envconfig:"MASSLOAD_ATTRIBUTE_SIZER"`
}

func WorkersDefault() Workers {
	return Workers{
		Page:           WorkerDefault(),
		Book:           WorkerDefault(),
		Hasher:         WorkerDefault(),
		Exporter:       WorkerDefault(),
		Tasker:         WorkerDefault(),
		FileValidator:  WorkerDefault(),
		FileTransferer: WorkerDefault(),
		MassloadSizer: Worker{
			Count:     1,
			QueueSize: 100,
			Interval:  time.Hour,
		},
		MassloadAttributeSizer: Worker{
			Count:     1,
			QueueSize: 100,
			Interval:  time.Hour,
		},
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

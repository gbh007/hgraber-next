package config

import "time"

type Workers struct {
	Page                   Worker `toml:"page" yaml:"page" envconfig:"PAGE"`
	Book                   Worker `toml:"book" yaml:"book" envconfig:"BOOK"`
	Hasher                 Worker `toml:"hasher" yaml:"hasher" envconfig:"HASHER"`
	Exporter               Worker `toml:"exporter" yaml:"exporter" envconfig:"EXPORTER"`
	Tasker                 Worker `toml:"tasker" yaml:"tasker" envconfig:"TASKER"`
	FileValidator          Worker `toml:"file_validator" yaml:"file_validator" envconfig:"FILE_VALIDATOR"`
	FileTransferer         Worker `toml:"file_transferer" yaml:"file_transferer" envconfig:"FILE_TRANSFERER"`
	MassloadSizer          Worker `toml:"massload_sizer" yaml:"massload_sizer" envconfig:"MASSLOAD_SIZER"`
	MassloadAttributeSizer Worker `toml:"massload_attribute_sizer" yaml:"massload_attribute_sizer" envconfig:"MASSLOAD_ATTRIBUTE_SIZER"`
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
	Count     int32         `toml:"count" yaml:"count" envconfig:"COUNT"`
	QueueSize int           `toml:"queue_size" yaml:"queue_size" envconfig:"QUEUE_SIZE"`
	Interval  time.Duration `toml:"interval" yaml:"interval" envconfig:"INTERVAL"`
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

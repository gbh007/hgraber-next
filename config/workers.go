package config

import (
	"time"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func WorkersDefault() []Worker {
	return []Worker{
		WorkerDefault(systemmodel.WorkerNamePage),
		WorkerDefault(systemmodel.WorkerNameBook),
		WorkerDefault(systemmodel.WorkerNameHasher),
		WorkerDefault(systemmodel.WorkerNameExporter),
		WorkerDefault(systemmodel.WorkerNameTasker),
		WorkerDefault(systemmodel.WorkerNameFileValidator),
		WorkerDefault(systemmodel.WorkerNameFileTransferer),
		{
			Name:      systemmodel.WorkerNameMassloadSizer,
			Count:     1,
			QueueSize: 100,
			Interval:  time.Hour,
		},
		{
			Name:      systemmodel.WorkerNameMassloadAttributeSizer,
			Count:     1,
			QueueSize: 100,
			Interval:  time.Hour,
		},
		{
			Name:      systemmodel.WorkerNameMassloadCalculation,
			Count:     1,
			QueueSize: 100,
			Interval:  time.Hour * 24,
		},
	}
}

type Worker struct {
	Name      string        `toml:"name" yaml:"name" envconfig:"NAME"`
	Count     int32         `toml:"count" yaml:"count" envconfig:"COUNT"`
	QueueSize int           `toml:"queue_size" yaml:"queue_size" envconfig:"QUEUE_SIZE"`
	Interval  time.Duration `toml:"interval" yaml:"interval" envconfig:"INTERVAL"`
}

func WorkerDefault(name string) Worker {
	return Worker{
		Name:      name,
		Count:     1,
		QueueSize: 100,
		Interval:  time.Minute,
	}
}

func (w Worker) GetName() string {
	return w.Name
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

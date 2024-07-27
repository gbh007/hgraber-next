package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgreSQLConnection  string        `envconfig:"POSTGRESQL_CONNECTION"`
	FilePath              string        `envconfig:"FILE_PATH"`
	WebServerAddr         string        `envconfig:"WEB_SERVER_ADDR"`
	ExternalWebServerAddr string        `envconfig:"EXTERNAL_WEB_SERVER_ADDR"`
	WebStaticDir          string        `envconfig:"WEB_STATIC_DIR"`
	APIToken              string        `envconfig:"API_TOKEN"`
	Debug                 bool          `envconfig:"DEBUG"`
	MetricTimeout         time.Duration `envconfig:"METRIC_TIMEOUT"`
	Handle                HandleConfig  `envconfig:"HANDLE"`
	FSAgentID             uuid.UUID     `envconfig:"FS_AGENT_ID"`
	Workers               Workers       `envconfig:"WORKERS"`
}

type HandleConfig struct {
	ParseBookTimeout time.Duration `envconfig:"PARSE_BOOK_TIMEOUT" default:"5m"`
}

type Workers struct {
	Page     Worker `envconfig:"PAGE"`
	Book     Worker `envconfig:"BOOK"`
	Hasher   Worker `envconfig:"HASHER"`
	Exporter Worker `envconfig:"EXPORTER"`
}

type Worker struct {
	Count     int32         `envconfig:"COUNT" default:"1"`
	QueueSize int           `envconfig:"QUEUE_SIZE" default:"100"`
	Interval  time.Duration `envconfig:"INTERVAL" default:"1m"`
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

func parseConfig() (Config, error) {
	c := Config{}

	err := envconfig.Process("APP", &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

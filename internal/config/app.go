package config

import "time"

type Application struct {
	Debug         bool          `yaml:"debug" envconfig:"DEBUG"`
	MetricTimeout time.Duration `yaml:"metric_timeout" envconfig:"METRIC_TIMEOUT"`
	ServiceName   string        `yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint string        `yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
}

func ApplicationDefault() Application {
	return Application{
		Debug:         false,
		MetricTimeout: 0,
		ServiceName:   "hgraber-next",
	}
}

package config

type Application struct {
	Metric        Metric    `toml:"metric" yaml:"metric" envconfig:"METRIC"`
	ServiceName   string    `toml:"service_name" yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint string    `toml:"trace_endpoint" yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	Pyroscope     Pyroscope `toml:"pyroscope" yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

type Pyroscope struct {
	Endpoint string `toml:"endpoint" yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `toml:"rate" yaml:"rate" envconfig:"RATE"`
}

func ApplicationDefault() Application {
	return Application{
		Metric:      MetricDefault(),
		ServiceName: "hgraber-next",
	}
}

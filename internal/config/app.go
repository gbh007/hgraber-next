package config

type ServerApplication struct {
	Metric        ServerMetric    `toml:"metric" yaml:"metric" envconfig:"METRIC"`
	ServiceName   string          `toml:"service_name" yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint string          `toml:"trace_endpoint" yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	Pyroscope     ServerPyroscope `toml:"pyroscope" yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

type ServerPyroscope struct {
	Endpoint string `toml:"endpoint" yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `toml:"rate" yaml:"rate" envconfig:"RATE"`
}

func ServerApplicationDefault() ServerApplication {
	return ServerApplication{
		Metric:      ServerMetricDefault(),
		ServiceName: "hgraber-next",
	}
}

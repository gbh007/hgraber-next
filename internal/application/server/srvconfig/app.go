package srvconfig

type ServerApplication struct {
	Metric      ServerMetric `toml:"metric"`
	ServiceName string       `toml:"service_name"`
}

func ServerApplicationDefault() ServerApplication {
	return ServerApplication{
		Metric:      ServerMetricDefault(),
		ServiceName: "hgraber-next",
	}
}

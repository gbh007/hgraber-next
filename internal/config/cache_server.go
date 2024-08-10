package config

// TODO: если не будет реализованно отдельное приложение, то удалить
type CacheServerApp struct {
	Addr        string `yaml:"addr" envconfig:"ADDR"`
	Token       string `yaml:"token" envconfig:"TOKEN"`
	TargetAddr  string `yaml:"target_addr" envconfig:"TARGET_ADDR"`
	TargetToken string `yaml:"target_token" envconfig:"TARGET_TOKEN"`

	Debug         bool   `yaml:"debug" envconfig:"DEBUG"`
	ServiceName   string `yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint string `yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
}

func CacheServerAppDefault() CacheServerApp {
	return CacheServerApp{
		Addr:        ":8080",
		ServiceName: "hgraber-next-cache-agent",
	}
}

type CacheServer struct {
	Addr  string `yaml:"addr" envconfig:"ADDR"`
	Token string `yaml:"token" envconfig:"TOKEN"`
}

func CacheServerDefault() CacheServer {
	return CacheServer{}
}

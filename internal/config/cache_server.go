package config

type CacheServer struct {
	Addr        string `yaml:"addr" envconfig:"ADDR"`
	Token       string `yaml:"token" envconfig:"TOKEN"`
	TargetAddr  string `yaml:"target_addr" envconfig:"TARGET_ADDR"`
	TargetToken string `yaml:"target_token" envconfig:"TARGET_TOKEN"`
}

func CacheServerDefault() CacheServer {
	return CacheServer{
		Addr:  ":8080",
		Token: "",
	}
}

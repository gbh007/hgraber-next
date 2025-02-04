package config

type AgentServer struct {
	Addr  string `yaml:"addr" envconfig:"ADDR"`
	Token string `yaml:"token" envconfig:"TOKEN"`
}

func AgentServerDefault() AgentServer {
	return AgentServer{}
}

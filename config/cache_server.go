package config

type AgentServer struct {
	Addr            string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	Token           string `toml:"token" yaml:"token" envconfig:"TOKEN"`
	LogErrorHandler bool   `toml:"log_error_handler" yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
}

func (a AgentServer) GetAddr() string {
	return a.Addr
}

func (a AgentServer) GetToken() string {
	return a.Token
}

func (a AgentServer) GetLogErrorHandler() bool {
	return a.LogErrorHandler
}

func (a AgentServer) GetDebug() bool {
	return a.Debug
}

func AgentServerDefault() AgentServer {
	return AgentServer{}
}

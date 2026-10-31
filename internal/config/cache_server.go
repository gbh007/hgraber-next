package config

type ServerAgentServer struct {
	Addr            string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	Token           string `toml:"token" yaml:"token" envconfig:"TOKEN"`
	LogErrorHandler bool   `toml:"log_error_handler" yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
}

func (a ServerAgentServer) GetAddr() string {
	return a.Addr
}

func (a ServerAgentServer) GetToken() string {
	return a.Token
}

func (a ServerAgentServer) GetLogErrorHandler() bool {
	return a.LogErrorHandler
}

func (a ServerAgentServer) GetDebug() bool {
	return a.Debug
}

func ServerAgentServerDefault() ServerAgentServer {
	return ServerAgentServer{}
}

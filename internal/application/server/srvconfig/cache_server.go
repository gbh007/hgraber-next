package srvconfig

type ServerAgentServer struct {
	Addr            string `toml:"addr"`
	Token           string `toml:"token"`
	LogErrorHandler bool   `toml:"log_error_handler"`
	Debug           bool   `toml:"debug"`
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

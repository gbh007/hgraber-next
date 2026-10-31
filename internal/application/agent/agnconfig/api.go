package agnconfig

type AgentAPI struct {
	Addr            string `toml:"addr"`
	Token           string `toml:"token"`
	LogErrorHandler bool   `toml:"log_error_handler"`
	Debug           bool   `toml:"debug"`
}

func (a AgentAPI) GetAddr() string {
	return a.Addr
}

func (a AgentAPI) GetToken() string {
	return a.Token
}

func (a AgentAPI) GetLogErrorHandler() bool {
	return a.LogErrorHandler
}

func (a AgentAPI) GetDebug() bool {
	return a.Debug
}

func AgentDefaultAPI() AgentAPI {
	return AgentAPI{
		Addr:  ":8080",
		Token: "",
	}
}

package config

type AgentAPI struct {
	Addr            string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	Token           string `toml:"token" yaml:"token" envconfig:"TOKEN"`
	LogErrorHandler bool   `toml:"log_error_handler" yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
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

package agnconfig

import "time"

type AgentDebugServer struct {
	Addr            string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	LogErrorHandler bool   `toml:"log_error_handler" yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
}

func AgentDefaultDebugServer() AgentDebugServer {
	return AgentDebugServer{}
}

type AgentParsers struct {
	HG4Token string            `toml:"hg4_token" yaml:"hg4_token" envconfig:"HG4_TOKEN"`
	Enabled  []string          `toml:"enabled" yaml:"enabled" envconfig:"ENABLED"`
	Cache    AgentParsersCache `toml:"cache" yaml:"cache" envconfig:"CACHE"`

	UserAgent string `toml:"user_agent" yaml:"user_agent" envconfig:"USER_AGENT"`
}

func AgentDefaultParsers() *AgentParsers {
	return &AgentParsers{
		HG4Token: "",
		Enabled: []string{
			"mock",
			"hgraber_local",
		},
		Cache: AgentDefaultParsersCache(),
	}
}

type AgentParsersCache struct {
	Enabled       bool          `toml:"enabled" yaml:"enabled" envconfig:"ENABLED"`
	Path          string        `toml:"path" yaml:"path" envconfig:"PATH"`
	TTL           time.Duration `toml:"ttl" yaml:"ttl" envconfig:"TTL"`
	CleanInterval time.Duration `toml:"clean_interval" yaml:"clean_interval" envconfig:"CLEAN_INTERVAL"`
}

func AgentDefaultParsersCache() AgentParsersCache {
	return AgentParsersCache{
		TTL:           time.Hour,
		CleanInterval: time.Minute * 5,
	}
}

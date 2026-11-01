package agnconfig

import "time"

type AgentDebugServer struct {
	Addr            string `toml:"addr"`
	LogErrorHandler bool   `toml:"log_error_handler"`
	Debug           bool   `toml:"debug"`
}

func AgentDefaultDebugServer() AgentDebugServer {
	return AgentDebugServer{}
}

type AgentParsers struct {
	Enabled   []string          `toml:"enabled"`
	Cache     AgentParsersCache `toml:"cache"`
	UserAgent string            `toml:"user_agent"`
	Proxy     AgentParsersProxy `toml:"proxy"`
}

type AgentParsersProxy struct {
	Host    string   `toml:"host"`
	User    string   `toml:"user"`
	Pass    string   `toml:"pass"`
	Enabled []string `toml:"enabled"`
}

func AgentDefaultParsers() AgentParsers {
	return AgentParsers{
		Enabled: []string{
			"mock",
			"hgraber_local",
		},
		Cache: AgentDefaultParsersCache(),
	}
}

type AgentParsersCache struct {
	Enabled       bool          `toml:"enabled"`
	Path          string        `toml:"path"`
	TTL           time.Duration `toml:"ttl"`
	CleanInterval time.Duration `toml:"clean_interval"`
}

func AgentDefaultParsersCache() AgentParsersCache {
	return AgentParsersCache{
		TTL:           time.Hour,
		CleanInterval: time.Minute * 5,
	}
}

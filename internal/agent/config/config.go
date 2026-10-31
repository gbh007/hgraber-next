package config

import "time"

type AgentConfig[T any] struct {
	Log         AgentLog         `toml:"log" yaml:"log" envconfig:"LOG"`
	API         AgentAPI         `toml:"api" yaml:"api" envconfig:"API"`
	DebugServer AgentDebugServer `toml:"debug_server" yaml:"debug_server" envconfig:"DEBUG_SERVER"`
	Application AgentApplication `toml:"application" yaml:"application" envconfig:"APPLICATION"`
	Parsers     *T               `toml:"parsers" yaml:"parsers" envconfig:"PARSERS"`
	FSBase      AgentFSBase      `toml:"fs_base" yaml:"fs_base" envconfig:"FS_BASE"`
	Sqlite      AgentSqlite      `toml:"sqlite" yaml:"sqlite" envconfig:"SQLITE"`
	ZipScanner  AgentZipScanner  `toml:"zip_scanner" yaml:"zip_scanner" envconfig:"ZIP_SCANNER"`
	Highway     AgentHighway     `toml:"highway" yaml:"highway" envconfig:"HIGHWAY"`
}

func AgentDefaultConfig[T any](defaultParsers func() *T) AgentConfig[T] {
	return AgentConfig[T]{
		Log:         AgentLogDefault(),
		API:         AgentDefaultAPI(),
		DebugServer: AgentDefaultDebugServer(),
		Parsers:     defaultParsers(),
		Application: AgentDefaultApplication(),
		FSBase:      AgentDefaultFSBase(),
		Sqlite:      AgentDefaultSqlite(),
		Highway:     AgentDefaultHighway(),
	}
}

func AgentDefaultConfigWrapped[T any](defaultParsers func() *T) func() AgentConfig[T] {
	return func() AgentConfig[T] {
		return AgentDefaultConfig(defaultParsers)
	}
}

type AgentApplication struct {
	TraceEndpoint   string         `toml:"trace_endpoint" yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	ClientTimeout   time.Duration  `toml:"client_timeout" yaml:"client_timeout" envconfig:"CLIENT_TIMEOUT"`
	ServiceName     string         `toml:"service_name" yaml:"service_name" envconfig:"SERVICE_NAME"`
	UseUnsafeCloser bool           `toml:"use_unsafe_closer" yaml:"use_unsafe_closer" envconfig:"USE_UNSAFE_CLOSER"`
	Pyroscope       AgentPyroscope `toml:"pyroscope" yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

type AgentPyroscope struct {
	Endpoint string `toml:"endpoint" yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `toml:"rate" yaml:"rate" envconfig:"RATE"`
}

func AgentDefaultApplication() AgentApplication {
	return AgentApplication{
		TraceEndpoint: "",
		ClientTimeout: time.Minute,
		ServiceName:   "hgraber-next-agent",
	}
}

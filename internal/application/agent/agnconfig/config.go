package agnconfig

import "time"

type AgentConfig struct {
	Log         AgentLog         `toml:"log"`
	API         AgentAPI         `toml:"api"`
	DebugServer AgentDebugServer `toml:"debug_server"`
	Application AgentApplication `toml:"application"`
	Parsers     AgentParsers     `toml:"parsers"`
	FSBase      AgentFSBase      `toml:"fs_base"`
	Sqlite      AgentSqlite      `toml:"sqlite"`
	ZipScanner  AgentZipScanner  `toml:"zip_scanner"`
	Highway     AgentHighway     `toml:"highway"`
}

func AgentDefaultConfig() AgentConfig {
	return AgentConfig{
		Log:         AgentLogDefault(),
		API:         AgentDefaultAPI(),
		DebugServer: AgentDefaultDebugServer(),
		Parsers:     AgentDefaultParsers(),
		Application: AgentDefaultApplication(),
		FSBase:      AgentDefaultFSBase(),
		Sqlite:      AgentDefaultSqlite(),
		Highway:     AgentDefaultHighway(),
	}
}

type AgentApplication struct {
	ClientTimeout time.Duration `toml:"client_timeout"`
	ServiceName   string        `toml:"service_name"`
}

func AgentDefaultApplication() AgentApplication {
	return AgentApplication{
		ClientTimeout: time.Minute,
		ServiceName:   "hgraber-next-agent",
	}
}

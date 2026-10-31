package config

type ServerStorage struct {
	DebugPGX       bool   `toml:"debug_pgx" yaml:"debug_pgx" envconfig:"DEBUG_PGX"`
	Connection     string `toml:"connection" yaml:"connection" envconfig:"CONNECTION"`
	MaxConnections int32  `toml:"max_connections" yaml:"max_connections" envconfig:"MAX_CONNECTIONS"`
}

type ServerFileStorage struct {
	TryReconnect bool `toml:"try_reconnect" yaml:"try_reconnect" envconfig:"TRY_RECONNECT"`
}

func ServerStorageDefault() ServerStorage {
	return ServerStorage{}
}

func ServerFileStorageDefault() ServerFileStorage {
	return ServerFileStorage{}
}

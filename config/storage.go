package config

type Storage struct {
	DebugSquirrel  bool   `toml:"debug_squirrel" yaml:"debug_squirrel" envconfig:"DEBUG_SQUIRREL"`
	DebugPGX       bool   `toml:"debug_pgx" yaml:"debug_pgx" envconfig:"DEBUG_PGX"`
	Connection     string `toml:"connection" yaml:"connection" envconfig:"CONNECTION"`
	MaxConnections int32  `toml:"max_connections" yaml:"max_connections" envconfig:"MAX_CONNECTIONS"`
}

func StorageDefault() Storage {
	return Storage{}
}

type FileStorage struct {
	TryReconnect bool `toml:"try_reconnect" yaml:"try_reconnect" envconfig:"TRY_RECONNECT"`
}

func FileStorageDefault() FileStorage {
	return FileStorage{}
}

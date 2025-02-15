package config

type Storage struct {
	Connection     string `yaml:"connection" envconfig:"CONNECTION"`
	MaxConnections int32  `yaml:"max_connections" envconfig:"MAX_CONNECTIONS"`
}

func StorageDefault() Storage {
	return Storage{}
}

type FileStorage struct {
	TryReconnect bool `yaml:"try_reconnect" envconfig:"TRY_RECONNECT"`
}

func FileStorageDefault() FileStorage {
	return FileStorage{}
}

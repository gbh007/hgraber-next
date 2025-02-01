package config

import "github.com/google/uuid"

type Storage struct {
	Connection     string    `yaml:"connection" envconfig:"CONNECTION"`
	MaxConnections int32     `yaml:"max_connections" envconfig:"MAX_CONNECTIONS"`
	FilePath       string    `yaml:"file_path" envconfig:"FILE_PATH"`
	FSAgentID      uuid.UUID `yaml:"fs_agent_id" envconfig:"FS_AGENT_ID"`
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

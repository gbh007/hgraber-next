package config

import "github.com/google/uuid"

type Storage struct {
	Connection string    `yaml:"connection" envconfig:"CONNECTION"`
	FilePath   string    `yaml:"file_path" envconfig:"FILE_PATH"`
	FSAgentID  uuid.UUID `yaml:"fs_agent_id" envconfig:"FS_AGENT_ID"`
}

func StorageDefault() Storage {
	return Storage{}
}

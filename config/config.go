package config

import (
	"time"
)

type Config struct {
	Log            Log            `yaml:"log" envconfig:"LOG"`
	Application    Application    `yaml:"application" envconfig:"APPLICATION"`
	Parsing        Parsing        `yaml:"parsing" envconfig:"PARSING"`
	Workers        Workers        `yaml:"workers" envconfig:"WORKERS"`
	Storage        Storage        `yaml:"storage" envconfig:"STORAGE"`
	FileStorage    FileStorage    `yaml:"file_storage" envconfig:"FILE_STORAGE"`
	API            API            `yaml:"api" envconfig:"API"`
	AgentServer    AgentServer    `yaml:"agent_server" envconfig:"AGENT_SERVER"`
	AttributeRemap AttributeRemap `yaml:"attribute_remap" envconfig:"ATTRIBUTE_REMAP"`
}

func ConfigDefault() Config {
	return Config{
		Log:         LogDefault(),
		Application: ApplicationDefault(),
		Parsing:     ParsingDefault(),
		API:         APIDefault(),
		Workers:     WorkersDefault(),
		Storage:     StorageDefault(),
		FileStorage: FileStorageDefault(),
		AgentServer: AgentServerDefault(),
	}
}

type Parsing struct {
	ParseBookTimeout time.Duration `yaml:"parse_book_timeout" envconfig:"PARSE_BOOK_TIMEOUT"`
	AgentTimeout     time.Duration `yaml:"agent_timeout" envconfig:"AGENT_TIMEOUT"`
}

func ParsingDefault() Parsing {
	return Parsing{
		ParseBookTimeout: time.Minute * 5,
		AgentTimeout:     time.Minute * 10,
	}
}

type AttributeRemap struct {
	Auto     bool `yaml:"auto" envconfig:"AUTO"`
	AllLower bool `yaml:"all_lower" envconfig:"ALL_LOWER"`
}

func AttributeRemapDefault() AttributeRemap {
	return AttributeRemap{}
}

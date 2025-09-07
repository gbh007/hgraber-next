package config

import (
	"time"
)

type Config struct {
	Log            Log            `toml:"log" yaml:"log" envconfig:"LOG"`
	Application    Application    `toml:"application" yaml:"application" envconfig:"APPLICATION"`
	Parsing        Parsing        `toml:"parsing" yaml:"parsing" envconfig:"PARSING"`
	Workers        []Worker       `toml:"workers" yaml:"workers" envconfig:"WORKERS"`
	Storage        Storage        `toml:"storage" yaml:"storage" envconfig:"STORAGE"`
	FileStorage    FileStorage    `toml:"file_storage" yaml:"file_storage" envconfig:"FILE_STORAGE"`
	API            API            `toml:"api" yaml:"api" envconfig:"API"`
	AgentServer    AgentServer    `toml:"agent_server" yaml:"agent_server" envconfig:"AGENT_SERVER"`
	AttributeRemap AttributeRemap `toml:"attribute_remap" yaml:"attribute_remap" envconfig:"ATTRIBUTE_REMAP"`
}

type Parsing struct {
	ParseBookTimeout time.Duration `toml:"parse_book_timeout" yaml:"parse_book_timeout" envconfig:"PARSE_BOOK_TIMEOUT"`
	AgentTimeout     time.Duration `toml:"agent_timeout" yaml:"agent_timeout" envconfig:"AGENT_TIMEOUT"`
}

type AttributeRemap struct {
	Auto     bool `toml:"auto" yaml:"auto" envconfig:"AUTO"`
	AllLower bool `toml:"all_lower" yaml:"all_lower" envconfig:"ALL_LOWER"`
}

func ConfigDefault() Config {
	return Config{
		Log:         LogDefault(),
		Application: ApplicationDefault(),
		Parsing:     ParsingDefault(),
		API:         APIDefault(),
		Storage:     StorageDefault(),
		FileStorage: FileStorageDefault(),
		AgentServer: AgentServerDefault(),
	}
}

func ParsingDefault() Parsing {
	return Parsing{
		ParseBookTimeout: DefaultParsingBookTimeout,
		AgentTimeout:     DefaultParsingAgentTimeout,
	}
}

func AttributeRemapDefault() AttributeRemap {
	return AttributeRemap{}
}

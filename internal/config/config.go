package config

import (
	"time"
)

type Config struct {
	Application Application `yaml:"application" envconfig:"APPLICATION"`
	Parsing     Parsing     `yaml:"parsing" envconfig:"PARSING"`
	Workers     Workers     `yaml:"workers" envconfig:"WORKERS"`
	Storage     Storage     `yaml:"storage" envconfig:"STORAGE"`
	API         API         `yaml:"api" envconfig:"API"`
	AgentServer AgentServer `yaml:"agent_server" envconfig:"AGENT_SERVER"`
}

func ConfigDefault() Config {
	return Config{
		Application: ApplicationDefault(),
		Parsing:     ParsingDefault(),
		API:         APIDefault(),
		Workers:     WorkersDefault(),
		Storage:     StorageDefault(),
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

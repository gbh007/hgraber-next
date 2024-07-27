package config

import (
	"time"
)

type Config struct {
	Debug         bool          `yaml:"debug" envconfig:"DEBUG"`
	MetricTimeout time.Duration `yaml:"metric_timeout" envconfig:"METRIC_TIMEOUT"`
	Parsing       Parsing       `yaml:"parsing" envconfig:"PARSING"`
	Workers       Workers       `yaml:"workers" envconfig:"WORKERS"`
	Storage       Storage       `yaml:"storage" envconfig:"STORAGE"`
	API           API           `yaml:"api" envconfig:"API"`
}

func ConfigDefault() Config {
	return Config{
		Parsing:       ParsingDefault(),
		Debug:         false,
		MetricTimeout: 0,
		API:           APIDefault(),
		Workers:       WorkersDefault(),
		Storage:       StorageDefault(),
	}
}

type Parsing struct {
	ParseBookTimeout time.Duration `yaml:"parse_book_timeout" envconfig:"PARSE_BOOK_TIMEOUT"`
}

func ParsingDefault() Parsing {
	return Parsing{
		ParseBookTimeout: time.Minute * 5,
	}
}

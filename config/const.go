package config

import "time"

const (
	DefaultParsingBookTimeout  = time.Minute * 5
	DefaultParsingAgentTimeout = time.Minute * 10

	DefaultMEtricScrapePeriodMainInfo      = 10 * time.Second
	DefaultMEtricScrapePeriodBookStatistic = time.Hour

	YamlIndent    = 2
	ConfigExtYaml = ".yaml"
	ConfigExtYml  = ".yml"
	ConfigExtToml = ".toml"
	ConfigExtEnv  = ".env"
)

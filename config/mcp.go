package config

type MCPServer struct {
	Addr   string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	Token  string `toml:"token" yaml:"token" envconfig:"TOKEN"`
	Debug  bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
	Mutate bool   `toml:"mutate" yaml:"mutate" envconfig:"MUTATE"`
}

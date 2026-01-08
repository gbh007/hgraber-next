package config

type MCPServer struct {
	Addr string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
}

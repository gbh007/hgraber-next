package srvconfig

type ServerMCPServer struct {
	Addr   string `toml:"addr"`
	Token  string `toml:"token"`
	Debug  bool   `toml:"debug"`
	Mutate bool   `toml:"mutate"`
}

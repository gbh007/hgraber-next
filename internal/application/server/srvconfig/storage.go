package srvconfig

type ServerStorage struct {
	DebugPGX       bool   `toml:"debug_pgx"`
	Connection     string `toml:"connection"`
	MaxConnections int32  `toml:"max_connections"`
}

type ServerFileStorage struct {
	TryReconnect bool `toml:"try_reconnect"`
}

func ServerStorageDefault() ServerStorage {
	return ServerStorage{}
}

func ServerFileStorageDefault() ServerFileStorage {
	return ServerFileStorage{}
}

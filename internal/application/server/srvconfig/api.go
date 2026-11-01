package srvconfig

import "os"

type ServerAPI struct {
	Addr                  string `toml:"addr"`
	ExternalAddr          string `toml:"external_addr"`
	StaticDir             string `toml:"static_dir"`
	Token                 string `toml:"token"`
	LogErrorHandler       bool   `toml:"log_error_handler"`
	Debug                 bool   `toml:"debug"`
	UseHeaderExternalAddr bool   `toml:"use_header_external_addr"`
}

func (a ServerAPI) GetAddr() string {
	return a.Addr
}

func (a ServerAPI) GetExternalAddr() string {
	return a.ExternalAddr
}

func (a ServerAPI) GetStaticDir() string {
	return a.StaticDir
}

func (a ServerAPI) GetToken() string {
	return a.Token
}

func (a ServerAPI) GetLogErrorHandler() bool {
	return a.LogErrorHandler
}

func (a ServerAPI) GetDebug() bool {
	return a.Debug
}

func (a ServerAPI) GetUseHeaderExternalAddr() bool {
	return a.UseHeaderExternalAddr
}

func ServerAPIDefault() ServerAPI {
	return ServerAPI{
		Addr:         ":8080",
		ExternalAddr: "http://localhost:8080",
		StaticDir:    os.Getenv("APP_API_STATIC_DIR"), // Для совместимости
		Token:        "",
	}
}

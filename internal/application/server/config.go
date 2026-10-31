package server

import (
	"flag"

	"github.com/gbh007/hgraber-next/external/config"
	"github.com/gbh007/hgraber-next/internal/application/server/srvconfig"
)

func parseConfig() (srvconfig.ServerConfig, error) {
	configPath := flag.String("config", "config.yaml", "path to config")

	flag.Parse()

	//nolint:wrapcheck // здесь обертка не нужна
	return config.ImportConfig(*configPath, true, srvconfig.ServerConfigDefault)
}

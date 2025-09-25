package server

import (
	"flag"

	"github.com/gbh007/hgraber-next/config"
)

func parseConfig() (config.Config, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	//nolint:wrapcheck // здесь обертка не нужна
	return config.ImportConfig(*configPath, true, config.ConfigDefault)
}

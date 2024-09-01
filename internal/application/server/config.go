package server

import (
	"flag"
	"os"

	"hgnext/internal/config"
)

func parseConfig() (config.Config, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	generateConfig := flag.String("generate-config", "", "generate example config")
	flag.Parse()

	if *generateConfig != "" {
		c := config.ConfigDefault()

		err := config.ExportToFile(&c, *generateConfig)
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	return config.ImportConfig(*configPath, true)
}

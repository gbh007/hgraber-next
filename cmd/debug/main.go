package main

import (
	"hgnext/internal/config"
)

func main() {
	serverConfig := config.ConfigDefault()
	agentConfig := config.CacheServerAppDefault()

	err := config.ExportToFile(&serverConfig, "example-server-config.yaml")
	if err != nil {
		panic(err)
	}

	err = config.ExportToFile(&agentConfig, "example-cache-agent-config.yaml")
	if err != nil {
		panic(err)
	}
}

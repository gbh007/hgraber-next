package main

import (
	"hgnext/internal/config"
)

func main() {
	serverConfig := config.ConfigDefault()

	err := config.ExportToFile(&serverConfig, "example-server-config.yaml")
	if err != nil {
		panic(err)
	}
}

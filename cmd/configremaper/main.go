package main

import (
	"github.com/gbh007/hgraber-next/internal/application/configremaper"
	"github.com/gbh007/hgraber-next/internal/config"
)

func main() {
	configremaper.Run(config.ServerConfigDefault)
}

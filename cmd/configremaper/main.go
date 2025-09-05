package main

import (
	"github.com/gbh007/hgraber-next/application/configremaper"
	"github.com/gbh007/hgraber-next/config"
)

func main() {
	configremaper.Run(config.ConfigDefault)
}

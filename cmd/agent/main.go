package main

import (
	"github.com/gbh007/hgraber-next/external/agentsdk"
	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parserloader/parsers/hgraber_local"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parserloader/parsers/mock"
)

func main() {
	agentsdk.InitAndServe(map[string]parserdto.NewParser{
		"mock":          mock.New,
		"hgraber_local": hgraber_local.New,
	})
}

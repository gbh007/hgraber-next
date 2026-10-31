package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
)

type Generator struct {
	uid             string
	services        []string
	useVictoriaLogs bool
}

func New(uid string, services []string, useVictoriaLogs bool) *Generator {
	plugins.RegisterDefaultPlugins()

	return &Generator{
		uid:             uid,
		services:        services,
		useVictoriaLogs: useVictoriaLogs,
	}
}

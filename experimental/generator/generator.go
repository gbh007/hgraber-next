package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
)

type Generator struct {
	uid      string
	services []string
}

func New(uid string, services []string) *Generator {
	plugins.RegisterDefaultPlugins()

	return &Generator{
		uid:      uid,
		services: services,
	}
}

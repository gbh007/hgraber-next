package logspannel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder, useVictoria bool) *dashboard.DashboardBuilder {
	builder.WithRow(dashboard.NewRowBuilder("Logs"))

	builder.WithPanel(generatorcore.WithPanelSize(Logs(useVictoria), generatorcore.PanelSizeFull))

	return builder
}

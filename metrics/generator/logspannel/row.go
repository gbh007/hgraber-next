package logspannel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(dashboard.NewRowBuilder("Logs"))

	builder.WithPanel(generatorcore.WithPanelSize(Logs(), generatorcore.PanelSizeFull))

	return builder
}

package databasepanel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(
		dashboard.NewRowBuilder("Database").
			WithPanel(generatorcore.WithPanelSize(LatencyQ95(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(RPS(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(ActiveRequest(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(OpenConnection(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(SlowRequest(), generatorcore.PanelSizeHalf)).
			Collapsed(true),
	)

	return builder
}

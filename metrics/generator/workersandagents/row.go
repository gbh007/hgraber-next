package workersandagents

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(
		dashboard.NewRowBuilder("Workers and agents").
			WithPanel(generatorcore.WithPanelSize(AgentParsingAvgLatency(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(AgentParsingRPS(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(WorkerAvgLatency(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(WorkerRPS(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(WorkerStat(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(WebCacheRPS(), generatorcore.PanelSizeHalf)).
			Collapsed(true),
	)

	return builder
}

package otherpanel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(
		dashboard.NewRowBuilder("Other").
			WithPanel(generatorcore.WithPanelSize(FSCompression(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(MetricCollectInfo(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(FSLatency(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(FSRPS(), generatorcore.PanelSizeHalf)).
			Collapsed(true),
	)

	return builder
}

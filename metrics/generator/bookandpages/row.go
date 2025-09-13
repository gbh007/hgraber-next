package bookandpages

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(
		dashboard.NewRowBuilder("Books and pages").
			WithPanel(generatorcore.WithPanelSize(BookCountDelta(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(PageCountDelta(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(FileCountDelta(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(FileSizeDelta(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(BookCount(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(PageCount(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(FileCount(), generatorcore.PanelSizeHalf)).
			WithPanel(generatorcore.WithPanelSize(FileSize(), generatorcore.PanelSizeHalf)).
			Collapsed(true),
	)

	return builder
}

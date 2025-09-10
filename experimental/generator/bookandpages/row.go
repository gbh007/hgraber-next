package bookandpages

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(
		dashboard.NewRowBuilder("Books and pages").
			WithPanel(generatorcore.WithPanelSize(BookCountDelta(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(PageCountDelta(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(FileCountDelta(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(FileSizeDelta(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(BookCount(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(PageCount(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(FileCount(), generatorcore.PanelSizeCommon)).
			WithPanel(generatorcore.WithPanelSize(FileSize(), generatorcore.PanelSizeCommon)).
			Collapsed(true),
	)

	return builder
}

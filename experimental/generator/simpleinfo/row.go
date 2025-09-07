package simpleinfo

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(dashboard.NewRowBuilder("Simple info"))

	builder.WithPanel(generatorcore.WithPanelSize(BookCount(), generatorcore.PanelSizeSlim))
	builder.WithPanel(generatorcore.WithPanelSize(PageCount(), generatorcore.PanelSizeSlim))
	builder.WithPanel(generatorcore.WithPanelSize(FileSize(), generatorcore.PanelSizeSlim))
	builder.WithPanel(generatorcore.WithPanelSize(FSCompression(), generatorcore.PanelSizeSlim))

	builder.WithPanel(generatorcore.WithPanelSize(BookCountDelta(), generatorcore.PanelSizeHalf))
	builder.WithPanel(generatorcore.WithPanelSize(PageCountDelta(), generatorcore.PanelSizeHalf))
	builder.WithPanel(generatorcore.WithPanelSize(FileCountDelta(), generatorcore.PanelSizeHalf))
	builder.WithPanel(generatorcore.WithPanelSize(FileSizeDelta(), generatorcore.PanelSizeHalf))

	return builder
}

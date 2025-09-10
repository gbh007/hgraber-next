package simpleinfo

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(dashboard.NewRowBuilder("Simple info"))

	builder.WithPanel(generatorcore.WithPanelSize(BookCount(), generatorcore.PanelSizeQuarterSlim))
	builder.WithPanel(generatorcore.WithPanelSize(PageCount(), generatorcore.PanelSizeQuarterSlim))
	builder.WithPanel(generatorcore.WithPanelSize(FileSize(), generatorcore.PanelSizeQuarterSlim))
	builder.WithPanel(generatorcore.WithPanelSize(FSCompression(), generatorcore.PanelSizeQuarterSlim))

	builder.WithPanel(generatorcore.WithPanelSize(BookCountDelta(), generatorcore.PanelSizeQuarter))
	builder.WithPanel(generatorcore.WithPanelSize(PageCountDelta(), generatorcore.PanelSizeQuarter))
	builder.WithPanel(generatorcore.WithPanelSize(FileCountDelta(), generatorcore.PanelSizeQuarter))
	builder.WithPanel(generatorcore.WithPanelSize(FileSizeDelta(), generatorcore.PanelSizeQuarter))

	builder.WithPanel(generatorcore.WithPanelSize(FSLatency(), generatorcore.PanelSizeHalf))
	builder.WithPanel(generatorcore.WithPanelSize(FSRPS(), generatorcore.PanelSizeHalf))

	return builder
}

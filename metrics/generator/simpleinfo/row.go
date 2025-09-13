package simpleinfo

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
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

	builder.WithPanel(generatorcore.WithPanelSize(WorkerErrorsDelta(), generatorcore.PanelSizeQuarterHigh))
	builder.WithPanel(generatorcore.WithPanelSize(FSSize(), generatorcore.PanelSizeQuarterHigh))
	builder.WithPanel(generatorcore.WithPanelSize(WorkerHandleDelta(), generatorcore.PanelSizeQuarterHigh))
	builder.WithPanel(generatorcore.WithPanelSize(ParserHandleDelta(), generatorcore.PanelSizeQuarterHigh))

	return builder
}

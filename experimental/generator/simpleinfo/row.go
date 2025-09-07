package simpleinfo

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
)

func WithRow(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithRow(dashboard.NewRowBuilder("Simple info"))
	builder.WithPanel(generatorcore.WithPanelSize(BookCount(), generatorcore.PanelSizeSlim))
	builder.WithPanel(generatorcore.WithPanelSize(PageCount(), generatorcore.PanelSizeSlim))

	return builder
}

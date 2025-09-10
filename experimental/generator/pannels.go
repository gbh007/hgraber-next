package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/experimental/generator/bookandpages"
	"github.com/gbh007/hgraber-next/experimental/generator/logspannel"
	"github.com/gbh007/hgraber-next/experimental/generator/simpleinfo"
	"github.com/gbh007/hgraber-next/experimental/generator/statistic"
)

func (g Generator) WithPanels(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	simpleinfo.WithRow(builder)
	logspannel.WithRow(builder)
	bookandpages.WithRow(builder)
	statistic.WithRow(builder)

	return builder
}

package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/metrics/generator/bookandpages"
	"github.com/gbh007/hgraber-next/metrics/generator/databasepanel"
	"github.com/gbh007/hgraber-next/metrics/generator/httpserverpanel"
	"github.com/gbh007/hgraber-next/metrics/generator/logspannel"
	"github.com/gbh007/hgraber-next/metrics/generator/otherpanel"
	"github.com/gbh007/hgraber-next/metrics/generator/simpleinfo"
	"github.com/gbh007/hgraber-next/metrics/generator/statistic"
	"github.com/gbh007/hgraber-next/metrics/generator/workersandagents"
)

func (g Generator) WithPanels(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	simpleinfo.WithRow(builder)
	logspannel.WithRow(builder)
	bookandpages.WithRow(builder)
	statistic.WithRow(builder)
	workersandagents.WithRow(builder)
	otherpanel.WithRow(builder)
	httpserverpanel.WithRow(builder)
	databasepanel.WithRow(builder)

	return builder
}

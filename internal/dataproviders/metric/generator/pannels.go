package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/bookandpages"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/databasepanel"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/httpserverpanel"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/logspannel"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/otherpanel"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/simpleinfo"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/statistic"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/workersandagents"
)

func (g Generator) WithPanels(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	simpleinfo.WithRow(builder)
	logspannel.WithRow(builder, g.useVictoriaLogs)
	bookandpages.WithRow(builder)
	statistic.WithRow(builder)
	workersandagents.WithRow(builder)
	otherpanel.WithRow(builder)
	httpserverpanel.WithRow(builder)
	databasepanel.WithRow(builder)

	return builder
}

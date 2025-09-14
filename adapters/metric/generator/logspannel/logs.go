package logspannel

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/logs"
	"github.com/grafana/grafana-foundation-sdk/go/loki"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
)

func Logs() *logs.PanelBuilder {
	query := promql.
		Vector("").
		Labels(generatorcore.ServiceFilterPromQL)

	return logs.
		NewPanelBuilder().
		Title("Logs").
		Targets([]cog.Builder[variants.Dataquery]{
			loki.
				NewDataqueryBuilder().
				Expr(query.String()).
				Datasource(generatorcore.LogsDatasource),
		}).
		SortOrder(common.LogsSortOrderDescending).
		EnableLogDetails(true).
		Datasource(generatorcore.LogsDatasource)
}

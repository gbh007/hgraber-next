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

func Logs(useVictoria bool) *logs.PanelBuilder {
	query := promql.
		Vector("").
		Labels(generatorcore.ServiceFilterPromQL)

	if useVictoria {
		return logs.
			NewPanelBuilder().
			Title("Logs").
			Targets([]cog.Builder[variants.Dataquery]{
				loki. // Примечание по сигнатуре частично совпадает, т.ч. используем его.
					NewDataqueryBuilder().
					Expr(`service_name : (` + generatorcore.NameToVar(generatorcore.ServiceVariableName) + `)`),
			}).
			SortOrder(common.LogsSortOrderDescending).
			EnableLogDetails(true).
			Datasource(generatorcore.LogsVictoriaLogsDatasource)
	}

	return logs.
		NewPanelBuilder().
		Title("Logs").
		Targets([]cog.Builder[variants.Dataquery]{
			loki.
				NewDataqueryBuilder().
				Expr(query.String()).
				Datasource(generatorcore.LogsLokiDatasource),
		}).
		SortOrder(common.LogsSortOrderDescending).
		EnableLogDetails(true).
		Datasource(generatorcore.LogsLokiDatasource)
}

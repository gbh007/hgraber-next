package generatorcore

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/pkg"
)

type PromQLExpr struct {
	Query  string
	Legend string
}

func SimpleTSPanel(exprs []PromQLExpr, title, unit string) *timeseries.PanelBuilder {
	return timeseries.
		NewPanelBuilder().
		Title(title).
		Targets(pkg.Map(exprs, func(expr PromQLExpr) cog.Builder[variants.Dataquery] {
			legend := LegendAuto
			if expr.Legend != "" {
				legend = expr.Legend
			}

			return prometheus.
				NewDataqueryBuilder().
				Expr(expr.Query).
				LegendFormat(legend)
		})).
		Legend(SimpleLegend()).
		Unit(unit).
		Thresholds(GreenTrashHold()).
		Datasource(MetricDatasource)
}

func SimpleTablePanel(exprs []PromQLExpr, title, unit string) *table.PanelBuilder {
	return table.NewPanelBuilder().
		Title(title).
		Targets(pkg.Map(exprs, func(expr PromQLExpr) cog.Builder[variants.Dataquery] {
			legend := LegendAuto
			if expr.Legend != "" {
				legend = expr.Legend
			}

			return prometheus.
				NewDataqueryBuilder().
				Expr(expr.Query).
				Format(prometheus.PromQueryFormatTable).
				Instant().
				LegendFormat(legend)
		})).
		Unit(unit).
		SortBy([]cog.Builder[common.TableSortByFieldState]{
			common.
				NewTableSortByFieldStateBuilder().
				DisplayName("Value").
				Desc(true),
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{
				Id:    "custom.hidden",
				Value: true,
			},
		}).
		Datasource(MetricDatasource)
}

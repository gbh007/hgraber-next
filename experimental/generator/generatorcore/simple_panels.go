package generatorcore

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
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

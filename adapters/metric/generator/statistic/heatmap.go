package statistic

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/heatmap"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
)

func Heatmap(title, metric, unit string) *heatmap.PanelBuilder {
	return heatmap.
		NewPanelBuilder().
		Title(title).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(
					promql.
						Sum(
							promql.
								Vector(metric + "_bucket").
								Labels(generatorcore.ServiceFilterPromQL),
						).
						By([]string{generatorcore.LELabel}).
						String(),
				).
				Format(prometheus.PromQueryFormatHeatmap).
				Interval(generatorcore.NameToVar(generatorcore.DeltaVariableName)).
				LegendFormat(generatorcore.LegendAuto),
		}).
		Mode(common.TooltipDisplayModeSingle).
		YAxis(heatmap.NewYAxisConfigBuilder().Unit(unit)).
		CellValues(
			heatmap.
				NewCellValuesBuilder().
				Unit(generatorcore.UnitShort),
		).
		FilterValues(
			heatmap.NewFilterValueRangeBuilder().Le(1),
		).
		Color(
			heatmap.
				NewHeatmapColorOptionsBuilder().
				Scheme("Blues").
				Mode(heatmap.HeatmapColorModeScheme).
				Steps(64). //nolint:mnd // будет исправлено позднее
				Fill("semi-dark-blue"),
		).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

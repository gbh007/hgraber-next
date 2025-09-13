package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/piechart"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func WorkerErrorsDelta() *piechart.PanelBuilder {
	return piechart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`Worker errors at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(generatorcore.IncreaseIntervalExpr(
					metricserver.WorkerExecutionTaskSecondsName+"_count",
					[]string{metriccore.SuccessLabel},
				)).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metriccore.SuccessLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		PieType(piechart.PieChartTypeDonut).
		Unit(generatorcore.UnitShort).
		Legend(piechart.
			NewPieChartLegendOptionsBuilder().
			Placement(common.LegendPlacementRight).
			DisplayMode(common.LegendDisplayModeTable).
			Values([]piechart.PieChartLegendValues{
				piechart.PieChartLegendValuesValue,
				piechart.PieChartLegendValuesPercent,
			}).
			ShowLegend(true),
		).
		OverrideByName(metriccore.OkLabelValue, []dashboard.DynamicConfigValue{
			{
				Id: "color",
				Value: map[string]string{
					"fixedColor": "green",
					"mode":       "fixed",
				},
			},
		}).
		OverrideByName(metriccore.ErrorLabelValue, []dashboard.DynamicConfigValue{
			{
				Id: "color",
				Value: map[string]string{
					"fixedColor": "red",
					"mode":       "fixed",
				},
			},
		}).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

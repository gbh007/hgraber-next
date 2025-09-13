package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/piechart"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FSSize() *piechart.PanelBuilder {
	query := promql.Sum(
		promql.
			Vector(metricserver.FileBytesName).
			Labels(generatorcore.ServiceFilterPromQL).
			Label(metricserver.TypeLabel, metricserver.TypeLabelValueFS),
	).By([]string{metriccore.FSIDLabel})

	return piechart.
		NewPanelBuilder().
		Title("FS size").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metriccore.FSIDLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		PieType(piechart.PieChartTypeDonut).
		Unit(generatorcore.UnitBytes).
		Legend(piechart.
			NewPieChartLegendOptionsBuilder().
			Placement(common.LegendPlacementBottom).
			DisplayMode(common.LegendDisplayModeTable).
			Values([]piechart.PieChartLegendValues{
				piechart.PieChartLegendValuesValue,
				piechart.PieChartLegendValuesPercent,
			}).
			ShowLegend(true),
		).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

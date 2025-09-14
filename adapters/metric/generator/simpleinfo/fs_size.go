package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/piechart"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func FSSize() *piechart.PanelBuilder {
	query := promql.Sum(
		promql.
			Vector(metricfs.FileBytesName).
			Labels(generatorcore.ServiceFilterPromQL).
			Label(metriccore.ServiceTypeLabel, metriccore.ServiceTypeLabelValueServer).
			Label(metriccore.TypeLabel, metricfs.TypeLabelValueFS),
	).By([]string{metricfs.FSIDLabel})

	return piechart.
		NewPanelBuilder().
		Title("FS size").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metricfs.FSIDLabel)).
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

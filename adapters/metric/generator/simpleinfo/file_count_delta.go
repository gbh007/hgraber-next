package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func FileCountDelta() *barchart.PanelBuilder {
	query := promql.Sum(
		promql.Delta(
			promql.
				Vector(metricfs.FileTotalName).
				Labels(generatorcore.ServiceFilterPromQL).
				Label(metriccore.ServiceTypeLabel, metriccore.ServiceTypeLabelValueServer).
				Range(generatorcore.NameToVar(generatorcore.DeltaVariableName)),
		),
	).By([]string{metriccore.TypeLabel})

	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`File delta count at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metriccore.TypeLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func PageCountDelta() *barchart.PanelBuilder {
	query := promql.Sum(
		promql.Delta(
			promql.
				Vector(metricserver.PageTotalName).
				Labels(generatorcore.ServiceFilterPromQL).
				Range(generatorcore.NameToVar(generatorcore.DeltaVariableName)),
		),
	).By([]string{metricserver.TypeLabel})

	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`Page delta count at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metricserver.TypeLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

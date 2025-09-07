package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func PageCountDelta() *barchart.PanelBuilder {
	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`Page delta count at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`sum(delta(%s{%s}[%s])) by (%s)`,
					metricserver.PageTotalName,
					generatorcore.ServiceFilter,
					generatorcore.NameToVar(generatorcore.DeltaVariableName),
					metricserver.TypeLabel,
				)).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metricserver.TypeLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Datasource(generatorcore.MetricDatasource)
}

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

func FileSizeDelta() *barchart.PanelBuilder {
	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`File delta size at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`sum(delta(%s{%s="%s", %s}[%s]))`,
					metricserver.FileBytesName,
					metricserver.TypeLabel,
					metricserver.TypeLabelValueFS,
					generatorcore.ServiceFilter,
					generatorcore.NameToVar(generatorcore.DeltaVariableName),
				)).
				Instant().
				LegendFormat("На диске").
				Datasource(generatorcore.MetricDatasource),
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`sum(delta(%s{%s="%s", %s}[%s]))`,
					metricserver.FileBytesName,
					metricserver.TypeLabel,
					metricserver.TypeLabelValuePage,
					generatorcore.ServiceFilter,
					generatorcore.NameToVar(generatorcore.DeltaVariableName),
				)).
				Instant().
				LegendFormat("В страницах").
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitBytes).
		Datasource(generatorcore.MetricDatasource)
}

package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FileSizeDelta() *barchart.PanelBuilder {
	query := func(k, v string) string {
		return promql.Sum(
			promql.Delta(
				promql.
					Vector(metricserver.FileBytesName).
					Labels(generatorcore.ServiceFilterPromQL).
					Label(k, v).
					Range(generatorcore.NameToVar(generatorcore.DeltaVariableName)),
			),
		).By([]string{metricserver.TypeLabel}).String()
	}

	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`File delta size at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query(
					metricserver.TypeLabel,
					metricserver.TypeLabelValueFS,
				)).
				Instant().
				LegendFormat("На диске").
				Datasource(generatorcore.MetricDatasource),
			prometheus.
				NewDataqueryBuilder().
				Expr(query(
					metricserver.TypeLabel,
					metricserver.TypeLabelValuePage,
				)).
				Instant().
				LegendFormat("В страницах").
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitBytes).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

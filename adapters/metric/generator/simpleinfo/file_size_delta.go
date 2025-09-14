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

func FileSizeDelta() *barchart.PanelBuilder {
	query := func(k, v string) string {
		return promql.Sum(
			promql.Delta(
				promql.
					Vector(metricfs.FileBytesName).
					Labels(generatorcore.ServiceFilterPromQL).
					Label(k, v).
					Label(metriccore.ServiceTypeLabel, metriccore.ServiceTypeLabelValueServer).
					Range(generatorcore.NameToVar(generatorcore.DeltaVariableName)),
			),
		).By([]string{metriccore.TypeLabel}).String()
	}

	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`File delta size at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query(
					metriccore.TypeLabel,
					metricfs.TypeLabelValueFS,
				)).
				Instant().
				LegendFormat("На диске").
				Datasource(generatorcore.MetricDatasource),
			prometheus.
				NewDataqueryBuilder().
				Expr(query(
					metriccore.TypeLabel,
					metricfs.TypeLabelValuePage,
				)).
				Instant().
				LegendFormat("В страницах").
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitBytes).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

package simpleinfo

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func FileSize() *stat.PanelBuilder {
	query := func(k, v string) string {
		return promql.Sum(
			promql.
				Vector(metricfs.FileBytesName).
				Labels(generatorcore.ServiceFilterPromQL).
				Label(metriccore.ServiceTypeLabel, metriccore.ServiceTypeLabelValueServer).
				Label(k, v),
		).By([]string{metriccore.TypeLabel}).String()
	}

	return stat.
		NewPanelBuilder().
		Title("File size").
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

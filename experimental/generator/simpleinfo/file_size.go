package simpleinfo

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FileSize() *stat.PanelBuilder {
	query := func(k, v string) string {
		return promql.Sum(
			promql.
				Vector(metricserver.FileBytesName).
				Labels(generatorcore.ServiceFilterPromQL).
				Label(k, v),
		).By([]string{metricserver.TypeLabel}).String()
	}

	return stat.
		NewPanelBuilder().
		Title("File size").
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
		Thresholds(
			dashboard.
				NewThresholdsConfigBuilder().
				Steps(generatorcore.GreenSteps),
		).
		Datasource(generatorcore.MetricDatasource)
}

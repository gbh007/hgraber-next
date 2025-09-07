package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FileSize() *stat.PanelBuilder {
	return stat.
		NewPanelBuilder().
		Title("File size").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`sum(%s{%s="%s", %s})`,
					metricserver.FileBytesName,
					metricserver.TypeLabel,
					metricserver.TypeLabelValueFS,
					generatorcore.ServiceFilter,
				)).
				Instant().
				LegendFormat("На диске").
				Datasource(generatorcore.MetricDatasource),
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`sum(%s{%s="%s", %s})`,
					metricserver.FileBytesName,
					metricserver.TypeLabel,
					metricserver.TypeLabelValuePage,
					generatorcore.ServiceFilter,
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

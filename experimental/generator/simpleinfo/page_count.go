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

func PageCount() *stat.PanelBuilder {
	return stat.
		NewPanelBuilder().
		Title("Page count").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`sum(%s{%s}) by (%s)`,
					metricserver.PageTotalName,
					generatorcore.ServiceFilter,
					metricserver.TypeLabel,
				)).
				LegendFormat(fmt.Sprintf("{{%s}}", metricserver.TypeLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Thresholds(
			dashboard.
				NewThresholdsConfigBuilder().
				Steps(generatorcore.GreenSteps),
		).
		Datasource(generatorcore.MetricDatasource)
}

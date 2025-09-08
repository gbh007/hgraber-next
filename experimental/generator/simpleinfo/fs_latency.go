package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricagent"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FSLatency() *timeseries.PanelBuilder {
	query := func(metric string, by []string) string {
		return promql.Div(
			promql.Sum(
				promql.Rate(
					promql.
						Vector(metric+"_sum").
						Labels(generatorcore.ServiceFilterPromQL).
						Range(generatorcore.RateIntervalVar),
				),
			).By(by),
			promql.Sum(
				promql.Rate(
					promql.
						Vector(metric+"_count").
						Labels(generatorcore.ServiceFilterPromQL).
						Range(generatorcore.RateIntervalVar),
				),
			).By(by),
		).String()
	}

	return timeseries.
		NewPanelBuilder().
		Title("FS latency").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query(
					metricserver.FSActionSecondsName,
					[]string{metriccore.ActionLabel, metriccore.FSIDLabel},
				)).
				LegendFormat(fmt.Sprintf("server/{{%s}} -> {{%s}}", metriccore.ActionLabel, metriccore.FSIDLabel)).
				Datasource(generatorcore.MetricDatasource),
			prometheus.
				NewDataqueryBuilder().
				Expr(query(
					metricagent.FSActionSecondsName,
					[]string{metriccore.ActionLabel},
				)).
				LegendFormat(fmt.Sprintf("agent/{{%s}}", metriccore.ActionLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Legend(generatorcore.SimpleLegend()).
		Unit(generatorcore.UnitSecond).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

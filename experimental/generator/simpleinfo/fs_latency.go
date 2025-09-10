package simpleinfo

import (
	"fmt"

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

	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: query(
					metricserver.FSActionSecondsName,
					[]string{metriccore.ActionLabel, metriccore.FSIDLabel},
				),
				Legend: fmt.Sprintf("server/{{%s}} -> {{%s}}", metriccore.ActionLabel, metriccore.FSIDLabel),
			},
			{
				Query: query(
					metricagent.FSActionSecondsName,
					[]string{metriccore.ActionLabel},
				),
				Legend: fmt.Sprintf("agent/{{%s}}", metriccore.ActionLabel),
			},
		},
		"FS latency",
		generatorcore.UnitSecond,
	)
}

package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricagent"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

func AgentParsingAvgLatency() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.AvgSummary(
					metricagent.ParserActionSecondsName,
					[]string{metriccore.ActionLabel, metricagent.ParserNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricagent.ParserNameLabel, metriccore.ActionLabel),
			},
		},
		"Agent parsing avg latency",
		generatorcore.UnitSecond,
	)
}

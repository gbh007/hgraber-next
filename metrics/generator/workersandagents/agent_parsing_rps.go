package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricagent"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
)

func AgentParsingRPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metricagent.ParserActionSecondsName+"_count",
					[]string{metriccore.ActionLabel, metricagent.ParserNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricagent.ParserNameLabel, metriccore.ActionLabel),
			},
		},
		"Agent parsing RPS",
		generatorcore.UnitRPS,
	)
}

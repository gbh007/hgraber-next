package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricagent"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

func WebCacheRPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metricagent.WebCacheTotalName,
					[]string{metriccore.ActionLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metriccore.ActionLabel),
			},
		},
		"Web cache RPS",
		generatorcore.UnitRPS,
	)
}

package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricagent"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
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

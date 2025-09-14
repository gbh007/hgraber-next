package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func WorkerRPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metricserver.WorkerExecutionTaskSecondsName+"_count",
					[]string{metricserver.WorkerNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.WorkerNameLabel),
			},
		},
		"Worker RPS",
		generatorcore.UnitRPS,
	)
}

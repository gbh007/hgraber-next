package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func WorkerAvgLatency() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.AvgSummary(
					metricserver.WorkerExecutionTaskSecondsName,
					[]string{metricserver.WorkerNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.WorkerNameLabel),
			},
		},
		"Worker avg latency",
		generatorcore.UnitSecond,
	)
}

package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func WorkerStat() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricserver.WorkerTotalName,
					[]string{metricserver.WorkerNameLabel, metricserver.CounterLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricserver.WorkerNameLabel, metricserver.CounterLabel),
			},
		},
		"Worker stat",
		generatorcore.UnitShort,
	)
}

package httpserverpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metrichttp"
)

func RPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metrichttp.ServerHandleDurationName+"_count",
					[]string{
						metriccore.ServiceNameLabel,
						metrichttp.ServerAddrLabelName,
						metrichttp.OperationLabelName,
						metrichttp.StatusLabelName,
					},
				),
				Legend: fmt.Sprintf(
					"{{%s}} {{%s}}/{{%s}} -> {{%s}}",
					metriccore.ServiceNameLabel,
					metrichttp.ServerAddrLabelName,
					metrichttp.OperationLabelName,
					metrichttp.StatusLabelName,
				),
			},
		},
		"RPS",
		generatorcore.UnitRPS,
	)
}

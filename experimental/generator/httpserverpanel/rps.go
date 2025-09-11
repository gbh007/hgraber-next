package httpserverpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metrichttp"
)

func RPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metrichttp.ServerHandleDurationName+"_count",
					[]string{metrichttp.ServerAddrLabelName, metrichttp.OperationLabelName, metrichttp.StatusLabelName},
				),
				Legend: fmt.Sprintf(
					"{{%s}}/{{%s}} -> {{%s}}",
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

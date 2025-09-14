package httpserverpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metrichttp"
)

func LatencyQ95() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RateQuantile(
					metrichttp.ServerHandleDurationName,
					[]string{
						metriccore.ServiceNameLabel,
						metrichttp.ServerAddrLabelName,
						metrichttp.OperationLabelName,
						metrichttp.StatusLabelName,
					},
					generatorcore.Q95,
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
		"Latency Q95",
		generatorcore.UnitSecond,
	)
}

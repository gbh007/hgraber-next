package httpserverpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metrichttp"
)

func ActiveRequest() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metrichttp.ServerActiveHandlersTotalName,
					[]string{metrichttp.ServerAddrLabelName, metrichttp.OperationLabelName},
				),
				Legend: fmt.Sprintf("{{%s}}/{{%s}}", metrichttp.ServerAddrLabelName, metrichttp.OperationLabelName),
			},
		},
		"Active request",
		generatorcore.UnitShort,
	)
}

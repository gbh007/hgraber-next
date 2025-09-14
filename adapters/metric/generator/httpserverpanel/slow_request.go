package httpserverpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metrichttp"
)

func SlowRequest() *table.PanelBuilder {
	query := promql.Div(
		promql.Sum(
			promql.
				Vector(metrichttp.ServerHandleDurationName+"_sum").
				Labels(generatorcore.ServiceFilterPromQL),
		).
			By([]string{metrichttp.ServerAddrLabelName, metrichttp.OperationLabelName, metrichttp.StatusLabelName}),
		promql.Sum(
			promql.
				Vector(metrichttp.ServerHandleDurationName+"_count").
				Labels(generatorcore.ServiceFilterPromQL),
		).
			By([]string{metrichttp.ServerAddrLabelName, metrichttp.OperationLabelName, metrichttp.StatusLabelName}),
	)

	return generatorcore.SimpleTablePanel(
		[]generatorcore.PromQLExpr{
			{
				Query: query.String(),
				Legend: fmt.Sprintf(
					"{{%s}}/{{%s}} -> {{%s}}",
					metrichttp.ServerAddrLabelName,
					metrichttp.OperationLabelName,
					metrichttp.StatusLabelName,
				),
			},
		},
		"Slow request",
		generatorcore.UnitSecond,
	)
}

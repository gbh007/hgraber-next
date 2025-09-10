package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func PageCount() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricserver.PageTotalName,
					[]string{metricserver.TypeLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.TypeLabel),
			},
		},
		"Page count",
		generatorcore.UnitShort,
	)
}

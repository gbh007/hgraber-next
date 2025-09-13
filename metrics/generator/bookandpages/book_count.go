package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func BookCount() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricserver.BookTotalName,
					[]string{metricserver.TypeLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.TypeLabel),
			},
		},
		"Book count",
		generatorcore.UnitShort,
	)
}

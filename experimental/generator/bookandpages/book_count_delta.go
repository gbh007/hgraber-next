package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func BookCountDelta() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.DeltaExpr(
					metricserver.BookTotalName,
					[]string{metricserver.TypeLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.TypeLabel),
			},
		},
		"Book delta count",
		generatorcore.UnitShort,
	)
}

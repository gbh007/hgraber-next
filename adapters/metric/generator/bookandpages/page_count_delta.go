package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func PageCountDelta() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.DeltaExpr(
					metricserver.PageTotalName,
					[]string{metricserver.TypeLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.TypeLabel),
			},
		},
		"Page delta count",
		generatorcore.UnitShort,
	)
}

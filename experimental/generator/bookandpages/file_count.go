package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FileCount() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricserver.FileTotalName,
					[]string{metricserver.TypeLabel, metriccore.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricserver.TypeLabel, metriccore.FSIDLabel),
			},
		},
		"File count",
		generatorcore.UnitShort,
	)
}

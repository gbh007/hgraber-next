package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FileSize() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricserver.FileBytesName,
					[]string{metricserver.TypeLabel, metriccore.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricserver.TypeLabel, metriccore.FSIDLabel),
			},
		},
		"File size",
		generatorcore.UnitBytes,
	)
}

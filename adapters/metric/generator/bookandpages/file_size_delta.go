package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func FileSizeDelta() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.DeltaExpr(
					metricserver.FileBytesName,
					[]string{metricserver.TypeLabel, metriccore.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricserver.TypeLabel, metriccore.FSIDLabel),
			},
		},
		"File delta size",
		generatorcore.UnitBytes,
	)
}

package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func FileSizeDelta() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.DeltaExpr(
					metricfs.FileBytesName,
					[]string{metriccore.TypeLabel, metricfs.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metriccore.TypeLabel, metricfs.FSIDLabel),
			},
		},
		"File delta size",
		generatorcore.UnitBytes,
	)
}

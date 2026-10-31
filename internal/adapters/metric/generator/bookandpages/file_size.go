package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/internal/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/internal/adapters/metric/metricfs"
)

func FileSize() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricfs.FileBytesName,
					[]string{metriccore.TypeLabel, metricfs.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metriccore.TypeLabel, metricfs.FSIDLabel),
			},
		},
		"File size",
		generatorcore.UnitBytes,
	)
}

package bookandpages

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricfs"
)

func FileCountDelta() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.DeltaExpr(
					metricfs.FileTotalName,
					[]string{metriccore.TypeLabel, metricfs.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metriccore.TypeLabel, metricfs.FSIDLabel),
			},
		},
		"File delta count",
		generatorcore.UnitShort,
	)
}

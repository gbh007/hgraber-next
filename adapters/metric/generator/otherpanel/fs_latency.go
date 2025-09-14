package otherpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func FSLatency() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.AvgSummary(
					metricfs.ActionSecondsName,
					[]string{metriccore.ServiceTypeLabel, metriccore.ActionLabel, metricfs.FSIDLabel},
				),
				Legend: fmt.Sprintf("{{%s}}/{{%s}} -> {{%s}}", metriccore.ServiceTypeLabel, metriccore.ActionLabel, metricfs.FSIDLabel),
			},
		},
		"FS latency",
		generatorcore.UnitSecond,
	)
}

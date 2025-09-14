package otherpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricagent"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func FSRPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metricserver.FSActionSecondsName+"_count",
					[]string{metriccore.ActionLabel, metriccore.FSIDLabel},
				),
				Legend: fmt.Sprintf("server/{{%s}} -> {{%s}}", metriccore.ActionLabel, metriccore.FSIDLabel),
			},
			{
				Query: generatorcore.RPSExpr(
					metricagent.FSActionSecondsName+"_count",
					[]string{metriccore.ActionLabel},
				),
				Legend: fmt.Sprintf("agent/{{%s}}", metriccore.ActionLabel),
			},
		},
		"FS RPS",
		generatorcore.UnitRPS,
	)
}

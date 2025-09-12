package databasepanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricdatabase"
)

func LatencyQ95() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RateQuantile(
					metricdatabase.RequestDurationName,
					[]string{metricdatabase.StmtLabelName},
					generatorcore.Q95,
				),
				Legend: fmt.Sprintf("{{%s}}", metricdatabase.StmtLabelName),
			},
		},
		"Latency Q95",
		generatorcore.UnitSecond,
	)
}

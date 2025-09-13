package databasepanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricdatabase"
)

func RPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metricdatabase.RequestDurationName+"_count",
					[]string{metricdatabase.StmtLabelName},
				),
				Legend: fmt.Sprintf("{{%s}}", metricdatabase.StmtLabelName),
			},
		},
		"RPS",
		generatorcore.UnitRPS,
	)
}

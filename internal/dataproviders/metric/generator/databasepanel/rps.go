package databasepanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricdatabase"
)

func RPS() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.RPSExpr(
					metricdatabase.RequestDurationName+"_count",
					[]string{metricdatabase.StmtLabelName, metricdatabase.DBLabelName},
				),
				Legend: fmt.Sprintf("{{%s}} -> {{%s}}", metricdatabase.DBLabelName, metricdatabase.StmtLabelName),
			},
		},
		"RPS",
		generatorcore.UnitRPS,
	)
}

package databasepanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricdatabase"
)

func OpenConnection() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricdatabase.OpenConnectionName,
					[]string{metricdatabase.DBLabelName},
				),
				Legend: fmt.Sprintf("{{%s}}", metricdatabase.DBLabelName),
			},
		},
		"Open connection",
		generatorcore.UnitShort,
	)
}

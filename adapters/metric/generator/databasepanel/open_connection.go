package databasepanel

import (
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
					nil,
				),
				Legend: "-",
			},
		},
		"Open connection",
		generatorcore.UnitShort,
	)
}

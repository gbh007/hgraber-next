package generatorcore

import "github.com/grafana/grafana-foundation-sdk/go/common"

func SimpleLegend() *common.VizLegendOptionsBuilder {
	return common.
		NewVizLegendOptionsBuilder().
		DisplayMode(common.LegendDisplayModeTable).
		Placement(common.LegendPlacementBottom).
		Calcs([]string{"mean", "lastNotNull"}).
		SortBy("Mean").
		SortDesc(true).
		ShowLegend(true)
}

func SimpleLegendLast() *common.VizLegendOptionsBuilder {
	return common.
		NewVizLegendOptionsBuilder().
		DisplayMode(common.LegendDisplayModeTable).
		Placement(common.LegendPlacementBottom).
		Calcs([]string{"lastNotNull"}).
		SortBy("Last *").
		SortDesc(true).
		ShowLegend(true)
}

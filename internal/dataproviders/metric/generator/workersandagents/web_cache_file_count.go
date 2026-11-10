package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricagent"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
)

func WebCacheFileCount() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricagent.WebCacheFileTotalName,
					[]string{metriccore.ServiceNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metriccore.ServiceNameLabel),
			},
		},
		"Web cache file count",
		generatorcore.UnitShort,
	)
}

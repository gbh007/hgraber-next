package workersandagents

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricagent"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
)

func WebCacheSize() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricagent.WebCacheBytesName,
					[]string{metriccore.ServiceNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metriccore.ServiceNameLabel),
			},
		},
		"Web cache size",
		generatorcore.UnitBytes,
	)
}

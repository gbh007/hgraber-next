package otherpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func MetricCollectInfo() *timeseries.PanelBuilder {
	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query: generatorcore.SumExpr(
					metricserver.LastCollectorScrapeDurationName,
					[]string{metricserver.ScrapeNameLabel},
				),
				Legend: fmt.Sprintf("{{%s}}", metricserver.ScrapeNameLabel),
			},
		},
		"Worker stat",
		generatorcore.UnitSecond,
	)
}

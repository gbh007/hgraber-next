package otherpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func FSCompression() *timeseries.PanelBuilder {
	query := promql.Sub(
		promql.N(1),
		promql.Div(
			promql.Sum(
				promql.
					Vector(metricserver.FileBytesName).
					Labels(generatorcore.ServiceFilterPromQL).
					Label(metricserver.TypeLabel, metricserver.TypeLabelValueFS),
			).By([]string{metriccore.FSIDLabel}),
			promql.Sum(
				promql.
					Vector(metricserver.FileBytesName).
					Labels(generatorcore.ServiceFilterPromQL).
					Label(metricserver.TypeLabel, metricserver.TypeLabelValuePage),
			).By([]string{metriccore.FSIDLabel}),
		),
	)

	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query:  query.String(),
				Legend: fmt.Sprintf("{{%s}}", metriccore.FSIDLabel),
			},
		},
		"FS Compression",
		generatorcore.UnitPercent0_1,
	)
}

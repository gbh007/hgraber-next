package otherpanel

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricfs"
)

func FSCompression() *timeseries.PanelBuilder {
	query := promql.Sub(
		promql.N(1),
		promql.Div(
			promql.Sum(
				promql.
					Vector(metricfs.FileBytesName).
					Labels(generatorcore.ServiceFilterPromQL).
					Label(metriccore.TypeLabel, metricfs.TypeLabelValueFS).
					Label(metriccore.ServiceTypeLabel, metriccore.ServiceTypeLabelValueServer),
			).By([]string{metricfs.FSIDLabel}),
			promql.Sum(
				promql.
					Vector(metricfs.FileBytesName).
					Labels(generatorcore.ServiceFilterPromQL).
					Label(metriccore.TypeLabel, metricfs.TypeLabelValuePage).
					Label(metriccore.ServiceTypeLabel, metriccore.ServiceTypeLabelValueServer),
			).By([]string{metricfs.FSIDLabel}),
		),
	)

	return generatorcore.SimpleTSPanel(
		[]generatorcore.PromQLExpr{
			{
				Query:  query.String(),
				Legend: fmt.Sprintf("{{%s}}", metricfs.FSIDLabel),
			},
		},
		"FS Compression",
		generatorcore.UnitPercent0_1,
	)
}

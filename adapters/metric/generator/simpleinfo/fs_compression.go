package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
)

func FSCompression() *stat.PanelBuilder {
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

	return stat.
		NewPanelBuilder().
		Title("FS Compression").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metriccore.FSIDLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitPercent0_1).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func BookCount() *stat.PanelBuilder {
	query := promql.Sum(
		promql.
			Vector(metricserver.BookTotalName).
			Labels(generatorcore.ServiceFilterPromQL),
	).By([]string{metricserver.TypeLabel})

	return stat.
		NewPanelBuilder().
		Title("Book count").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metricserver.TypeLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metricserver"
)

func PageCountDelta() *barchart.PanelBuilder {
	query := promql.Sum(
		promql.Delta(
			promql.
				Vector(metricserver.PageTotalName).
				Labels(generatorcore.ServiceFilterPromQL).
				Range(generatorcore.RangeIntervalVar),
		),
	).By([]string{metriccore.TypeLabel})

	return barchart.
		NewPanelBuilder().
		Title("Page delta count").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(query.String()).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metriccore.TypeLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

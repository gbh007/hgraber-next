package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func WorkerHandleDelta() *barchart.PanelBuilder {
	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`Worker handle at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(generatorcore.IncreaseIntervalExpr(
					metricserver.WorkerExecutionTaskSecondsName+"_count",
					[]string{metricserver.WorkerNameLabel, metriccore.SuccessLabel},
				)).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}/{{%s}}", metricserver.WorkerNameLabel, metriccore.SuccessLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Legend(generatorcore.SimpleLegendLast()).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

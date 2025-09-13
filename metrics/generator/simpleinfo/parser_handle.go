package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/barchart"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"

	"github.com/gbh007/hgraber-next/metrics/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricagent"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
)

func ParserHandleDelta() *barchart.PanelBuilder {
	return barchart.
		NewPanelBuilder().
		Title(fmt.Sprintf(`Parser handle at %s`, generatorcore.NameToVar(generatorcore.DeltaVariableName))).
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(generatorcore.IncreaseIntervalExpr(
					metricagent.ParserActionSecondsName+"_count",
					[]string{metriccore.ActionLabel, metricagent.ParserNameLabel},
				)).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}} -> {{%s}}", metricagent.ParserNameLabel, metriccore.ActionLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitShort).
		Legend(generatorcore.SimpleLegendLast()).
		Thresholds(generatorcore.GreenTrashHold()).
		Datasource(generatorcore.MetricDatasource)
}

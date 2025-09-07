package simpleinfo

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metriccore"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func FSCompression() *stat.PanelBuilder {
	return stat.
		NewPanelBuilder().
		Title("FS Compression").
		Targets([]cog.Builder[variants.Dataquery]{
			prometheus.
				NewDataqueryBuilder().
				Expr(fmt.Sprintf(
					`1 - 
              sum(%[1]s{%[2]s="%[5]s", %[3]s}) by (%[4]s)
              /
              sum(%[1]s{%[2]s="%[6]s", %[3]s}) by (%[4]s)`,
					metricserver.FileBytesName,
					metricserver.TypeLabel,
					generatorcore.ServiceFilter,
					metriccore.FSIDLabel,
					metricserver.TypeLabelValueFS,
					metricserver.TypeLabelValuePage,
				)).
				Instant().
				LegendFormat(fmt.Sprintf("{{%s}}", metriccore.FSIDLabel)).
				Datasource(generatorcore.MetricDatasource),
		}).
		Unit(generatorcore.UnitPercent0_1).
		Thresholds(
			dashboard.
				NewThresholdsConfigBuilder().
				Steps(generatorcore.GreenSteps),
		).
		Datasource(generatorcore.MetricDatasource)
}

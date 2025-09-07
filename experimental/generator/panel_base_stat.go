package generator

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"

	"github.com/gbh007/hgraber-next/experimental/generator/core"
)

func (g Generator) withStatPanels(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithPanel(
		core.WithPanelSize(
			stat.
				NewPanelBuilder().
				Title("Book count").
				Targets([]cog.Builder[variants.Dataquery]{
					prometheus.
						NewDataqueryBuilder().
						Expr(fmt.Sprintf(
							`sum(hgraber_next_server_book_total{%s}) by (type)`,
							core.ServiceFilter,
						)).
						LegendFormat("{{type}}").
						Datasource(core.MetricDatasource),
				}).
				Unit(core.UnitShort).
				Thresholds(
					dashboard.
						NewThresholdsConfigBuilder().
						Steps(core.GreenSteps),
				).
				Datasource(core.MetricDatasource),
			core.PanelSizeSlim,
		),
	)

	builder.WithPanel(
		core.WithPanelSize(
			stat.
				NewPanelBuilder().
				Title("Page count").
				Targets([]cog.Builder[variants.Dataquery]{
					prometheus.
						NewDataqueryBuilder().
						Expr(fmt.Sprintf(
							`sum(hgraber_next_server_page_total{%s}) by (type)`,
							core.ServiceFilter,
						)).
						LegendFormat("{{type}}").
						Datasource(core.MetricDatasource),
				}).
				Unit(core.UnitShort).
				Thresholds(
					dashboard.
						NewThresholdsConfigBuilder().
						Steps(core.GreenSteps),
				).
				Datasource(core.MetricDatasource),
			core.PanelSizeSlim,
		),
	)

	return builder
}

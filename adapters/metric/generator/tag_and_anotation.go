package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

func (g Generator) WithTagAndAnnotation(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	tags := []string{"hgnext"}

	builder.Tags(tags)

	builder.Link(
		dashboard.
			NewDashboardLinkBuilder("GitHub").
			Url("https://github.com/gbh007/hgraber-next").
			Type(dashboard.DashboardLinkTypeLink).
			TargetBlank(true),
	)

	builder.Link(
		dashboard.
			NewDashboardLinkBuilder("HG next boards").
			Tags(tags).
			Type(dashboard.DashboardLinkTypeDashboards).
			KeepTime(true).
			AsDropdown(true).
			TargetBlank(true),
	)

	if !g.useVictoriaLogs {
		builder.Annotation(
			dashboard.
				NewAnnotationQueryBuilder().
				Enable(true).
				// TODO: привести в более аккуратный вид
				Expr("{service_name=~\"$service_name\"} |= `application start`").
				IconColor("super-light-purple").
				Name("app started (logs)").
				Datasource(generatorcore.LogsLokiDatasource),
		)
	}

	builder.Annotation(
		dashboard.
			NewAnnotationQueryBuilder().
			Enable(true).
			Expr(
				promql.Sum(
					promql.Changes(
						promql.
							Vector(metriccore.VersionInfoName).
							Labels(generatorcore.ServiceFilterPromQL),
					),
				).
					By([]string{metriccore.ServiceNameLabel}).
					String(),
			).
			IconColor("super-light-blue").
			Placement(dashboard.AnnotationQueryPlacementInControlsMenu).
			Name("app started (metrics)").
			// TODO: перенастроить как будут изменения в либе графаны
			// Title(fmt.Sprintf("{{%s}}", metriccore.ServiceNameLabel)).
			Datasource(generatorcore.MetricDatasource),
	)

	return builder
}

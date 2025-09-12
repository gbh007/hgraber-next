package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/promql-builder/go/promql"

	"github.com/gbh007/hgraber-next/experimental/generator/generatorcore"
	"github.com/gbh007/hgraber-next/metrics/metricagent"
	"github.com/gbh007/hgraber-next/metrics/metricserver"
)

func (g Generator) WithTagAndAnotation(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
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

	builder.Annotation(
		dashboard.
			NewAnnotationQueryBuilder().
			Enable(true).
			Expr("{service_name=~\"$service_name\"} |= `application start`"). // TODO: привести в более аккуратный вид
			IconColor("super-light-purple").
			Name("app started (logs)").
			Datasource(generatorcore.LogsDatasource),
	)

	builder.Annotation( // TODO: перенастроить как будут изменения в либе графаны
		dashboard.
			NewAnnotationQueryBuilder().
			Enable(false).
			Expr(promql.Or(
				promql.Mul(
					promql.
						Vector(metricagent.VersionInfoName).
						Labels(generatorcore.ServiceFilterPromQL),
					promql.N(1000), //nolint:mnd // будет исправлено позднее
				),
				promql.Mul(
					promql.
						Vector(metricserver.VersionInfoName).
						Labels(generatorcore.ServiceFilterPromQL),
					promql.N(1000), //nolint:mnd // будет исправлено позднее
				),
			).String()).
			IconColor("super-light-blue").
			Name("app started (metrics)").
			Datasource(generatorcore.MetricDatasource),
	)

	return builder
}

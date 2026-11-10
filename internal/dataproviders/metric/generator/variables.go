package generator

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric/metriccore"
)

func (g Generator) WithVariables(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	if g.useVictoriaLogs {
		builder.WithVariable(
			dashboard.
				NewDatasourceVariableBuilder(generatorcore.LogsVariableName).
				Type(generatorcore.LogsVariableTypeVictoriaLogs),
		)
	} else {
		builder.WithVariable(
			dashboard.
				NewDatasourceVariableBuilder(generatorcore.LogsVariableName).
				Type(generatorcore.LogsVariableTypeLoki),
		)
	}

	builder.WithVariable(
		dashboard.
			NewDatasourceVariableBuilder(generatorcore.MetricVariableName).
			Type(generatorcore.MetricVariableType),
	)
	builder.WithVariable(
		dashboard.
			NewQueryVariableBuilder(generatorcore.ServiceVariableName).
			Query(generatorcore.ValuesFromString(fmt.Sprintf(
				"label_values(%s, %s)",
				metriccore.VersionInfoName,
				metriccore.ServiceNameLabel,
			))).
			Datasource(generatorcore.MetricDatasource).
			Multi(true).
			Current(dashboard.VariableOption{
				Selected: new(true),
				Text: dashboard.StringOrArrayOfString{
					ArrayOfString: g.services,
				},
				Value: dashboard.StringOrArrayOfString{
					ArrayOfString: g.services,
				},
			}).
			Refresh(dashboard.VariableRefreshOnTimeRangeChanged),
	)
	return builder
}

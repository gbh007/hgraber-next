package generator

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/adapters/metric/generator/generatorcore"
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

//nolint:lll // будет исправлено позднее
func (g Generator) WithVariables(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithVariable(
		dashboard.
			NewDatasourceVariableBuilder(generatorcore.LogsVariableName).
			Type(generatorcore.LogsVariableType),
	)
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
				Selected: generatorcore.BoolToPtr(true),
				Text: dashboard.StringOrArrayOfString{
					ArrayOfString: g.services,
				},
				Value: dashboard.StringOrArrayOfString{
					ArrayOfString: g.services,
				},
			}).
			Refresh(dashboard.VariableRefreshOnTimeRangeChanged),
	)
	builder.WithVariable(
		dashboard.
			NewIntervalVariableBuilder(generatorcore.DeltaVariableName).
			Values(generatorcore.ValuesFromArray(generatorcore.DeltaVariableValues)).
			Current(dashboard.VariableOption{
				Selected: generatorcore.BoolToPtr(true),
				Text: dashboard.StringOrArrayOfString{
					String: generatorcore.StrToPtr(generatorcore.DeltaVariableCurrent),
				},
				Value: dashboard.StringOrArrayOfString{
					String: generatorcore.StrToPtr(generatorcore.DeltaVariableCurrent),
				},
			}).
			Auto(true),
	)

	return builder
}

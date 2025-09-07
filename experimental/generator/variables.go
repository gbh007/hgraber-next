package generator

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"

	"github.com/gbh007/hgraber-next/experimental/generator/core"
)

//nolint:lll // будет исправлено позднее
func (g Generator) WithVariables(builder *dashboard.DashboardBuilder) *dashboard.DashboardBuilder {
	builder.WithVariable(
		dashboard.
			NewDatasourceVariableBuilder(core.LogsVariableName).
			Type(core.LogsVariableType),
	)
	builder.WithVariable(
		dashboard.
			NewDatasourceVariableBuilder(core.MetricVariableName).
			Type(core.MetricVariableType),
	)
	builder.WithVariable(
		dashboard.
			NewQueryVariableBuilder(core.ServiceVariableName).
			Query(core.ValuesFromString(`label_values({__name__=~ "hgraber_next_server_version_info|hgraber_next_agent_version_info"}, service_name)`)).
			Datasource(core.MetricDatasource).
			Multi(true).
			Current(dashboard.VariableOption{
				Selected: core.BoolToPtr(true),
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
			NewIntervalVariableBuilder(core.DeltaVariableName).
			Values(core.ValuesFromArray(core.DeltaVariableValues)).
			Current(dashboard.VariableOption{
				Selected: core.BoolToPtr(true),
				Text: dashboard.StringOrArrayOfString{
					String: core.StrToPtr(core.DeltaVariableCurrent),
				},
				Value: dashboard.StringOrArrayOfString{
					String: core.StrToPtr(core.DeltaVariableCurrent),
				},
			}).
			Auto(true),
	)

	return builder
}

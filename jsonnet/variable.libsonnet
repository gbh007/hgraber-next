local config = import 'config.libsonnet';
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

{
  local variable = grafonnet.dashboard.variable,
  metrics():
    variable.datasource.new(
      config.variable.prometheus.name,
      config.variable.prometheus.type,
    ),
  logs():
    variable.datasource.new(
      config.variable.loki.name,
      config.variable.loki.type,
    ),
  service():
    variable.query.new(
      config.variable.serviceName.name,
      'label_values({__name__=~ "hgraber_next_server_version_info|hgraber_next_agent_version_info"}, service_name)',
    )
    + variable.query.withDatasource(config.datasource.metrics.type, config.datasource.metrics.uid)
    + variable.query.selectionOptions.withMulti()
    + variable.query.generalOptions.withCurrent(config.variable.serviceName.current)
    + variable.query.refresh.onTime(),
  delta():
    variable.interval.new(config.variable.delta.name, config.variable.delta.values)
    + variable.interval.generalOptions.withCurrent(config.variable.delta.current)
    + variable.interval.withAutoOption(),
}

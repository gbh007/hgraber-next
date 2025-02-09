local config = import 'config.libsonnet';
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

local panel = grafonnet.panel;
local query = grafonnet.query;
local prometheus = grafonnet.query.prometheus;

local greenSteps() = { color: 'green', value: null };

{
  pannels: [
    panel.row.new('Simple info'),
    panel.stat.new('Book count')
    + panel.stat.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'hgraber_next_server_book_total{%s}' % config.label.filter.service,
      )
      + prometheus.withLegendFormat('{{type}}')
      + prometheus.withInstant(),
    ])
    + panel.stat.standardOptions.withUnit('short')
    + panel.stat.standardOptions.thresholds.withSteps([greenSteps()])
    + panel.stat.gridPos.withH(3)
    + panel.stat.gridPos.withW(12)
    + panel.stat.gridPos.withX(0)
    + panel.stat.gridPos.withY(1)
    + panel.stat.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.stat.new('File size')
    + panel.stat.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(hgraber_next_server_file_bytes{type="fs",%s})' % config.label.filter.service,
      )
      + prometheus.withLegendFormat('На диске')
      + prometheus.withInstant(),
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(hgraber_next_server_file_bytes{type="page",%s})' % config.label.filter.service,
      )
      + prometheus.withLegendFormat('В страницах')
      + prometheus.withInstant(),
    ])
    + panel.stat.standardOptions.withUnit('bytes')
    + panel.stat.standardOptions.thresholds.withSteps([greenSteps()])
    + panel.stat.gridPos.withH(3)
    + panel.stat.gridPos.withW(12)
    + panel.stat.gridPos.withX(12)
    + panel.stat.gridPos.withY(1)
    + panel.stat.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.barChart.new('Book delta count at $%s' % config.variable.delta.name)
    + panel.barChart.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(delta(hgraber_next_server_book_total{%s}[$%s])) by (type)' % [
          config.label.filter.service,
          config.variable.delta.name,
        ],
      )
      + prometheus.withLegendFormat('{{type}}')
      + prometheus.withInstant(),
    ])
    + panel.barChart.gridPos.withH(9)
    + panel.barChart.gridPos.withW(6)
    + panel.barChart.gridPos.withX(0)
    + panel.barChart.gridPos.withY(4)
    + panel.barChart.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.barChart.new('Page delta count at $%s' % config.variable.delta.name)
    + panel.barChart.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(delta(hgraber_next_server_page_total{%s}[$%s])) by (type)' % [
          config.label.filter.service,
          config.variable.delta.name,
        ],
      )
      + prometheus.withLegendFormat('{{type}}')
      + prometheus.withInstant(),
    ])
    + panel.barChart.gridPos.withH(9)
    + panel.barChart.gridPos.withW(6)
    + panel.barChart.gridPos.withX(6)
    + panel.barChart.gridPos.withY(4)
    + panel.barChart.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.barChart.new('File delta count at $%s' % config.variable.delta.name)
    + panel.barChart.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(delta(hgraber_next_server_file_total{%s}[$%s])) by (type)' % [
          config.label.filter.service,
          config.variable.delta.name,
        ],
      )
      + prometheus.withLegendFormat('{{type}}')
      + prometheus.withInstant(),
    ])
    + panel.barChart.gridPos.withH(9)
    + panel.barChart.gridPos.withW(6)
    + panel.barChart.gridPos.withX(12)
    + panel.barChart.gridPos.withY(4)
    + panel.barChart.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.barChart.new('File delta size at $%s' % config.variable.delta.name)
    + panel.barChart.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(delta(hgraber_next_server_file_bytes{type="fs", %s}[$%s]))' % [
          config.label.filter.service,
          config.variable.delta.name,
        ],
      )
      + prometheus.withLegendFormat('На диске')
      + prometheus.withInstant(),
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(delta(hgraber_next_server_file_bytes{type="page", %s}[$%s]))' % [
          config.label.filter.service,
          config.variable.delta.name,
        ],
      )
      + prometheus.withLegendFormat('В страницах')
      + prometheus.withInstant(),
    ])
    + panel.stat.standardOptions.withUnit('bytes')
    + panel.barChart.gridPos.withH(9)
    + panel.barChart.gridPos.withW(6)
    + panel.barChart.gridPos.withX(18)
    + panel.barChart.gridPos.withY(4)
    + panel.barChart.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.timeSeries.new('FS latency')
    + panel.timeSeries.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        |||
          sum(rate(hgraber_next_server_fs_action_seconds_sum{%s}[$__rate_interval])) by (action, fs_id)
          /
          sum(rate(hgraber_next_server_fs_action_seconds_count{%s}[$__rate_interval])) by (action, fs_id)
        |||
        % [
          config.label.filter.service,
          config.label.filter.service,
        ],
      )
      + prometheus.withLegendFormat('server/{{action}} -> {{fs_id}}'),
      prometheus.new(
        config.datasource.metrics.uid,
        |||
          sum(rate(hgraber_next_agent_fs_action_seconds_sum{%s}[$__rate_interval])) by (action)
          /
          sum(rate(hgraber_next_agent_fs_action_seconds_count{%s}[$__rate_interval])) by (action)
        |||
        % [
          config.label.filter.service,
          config.label.filter.service,
        ],
      )
      + prometheus.withLegendFormat('agent/{{action}}'),
    ])
    + panel.timeSeries.standardOptions.withUnit('s')
    + panel.timeSeries.options.legend.withDisplayMode('table')
    + panel.timeSeries.options.legend.withPlacement('bottom')
    + panel.timeSeries.options.legend.withCalcs(['mean', 'lastNotNull'])
    + panel.timeSeries.options.legend.withSortBy('mean')
    + panel.timeSeries.options.legend.withSortDesc()
    + panel.timeSeries.gridPos.withH(9)
    + panel.timeSeries.gridPos.withW(12)
    + panel.timeSeries.gridPos.withX(0)
    + panel.timeSeries.gridPos.withY(13)
    + panel.timeSeries.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
    panel.timeSeries.new('FS RPS')
    + panel.timeSeries.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(rate(hgraber_next_server_fs_action_seconds_count{%s}[$__rate_interval])) by (action, fs_id)' % [
          config.label.filter.service,
        ],
      )
      + prometheus.withLegendFormat('server/{{action}} -> {{fs_id}}'),
      prometheus.new(
        config.datasource.metrics.uid,
        'sum(rate(hgraber_next_agent_fs_action_seconds_count{%s}[$__rate_interval])) by (action)' % [
          config.label.filter.service,
        ],
      )
      + prometheus.withLegendFormat('agent/{{action}}'),
    ])
    + panel.timeSeries.standardOptions.withUnit('reqps')
    + panel.timeSeries.options.legend.withDisplayMode('table')
    + panel.timeSeries.options.legend.withPlacement('bottom')
    + panel.timeSeries.options.legend.withCalcs(['mean', 'lastNotNull'])
    + panel.timeSeries.options.legend.withSortBy('mean')
    + panel.timeSeries.options.legend.withSortDesc()
    + panel.timeSeries.gridPos.withH(9)
    + panel.timeSeries.gridPos.withW(12)
    + panel.timeSeries.gridPos.withX(12)
    + panel.timeSeries.gridPos.withY(13)
    + panel.timeSeries.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
  ],
}

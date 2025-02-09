local config = import 'config.libsonnet';
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

local panel = grafonnet.panel;
local query = grafonnet.query;
local prometheus = grafonnet.query.prometheus;

local greenSteps() = { color: 'green', value: null };
local simpleTSLegend() =
  panel.timeSeries.options.legend.withDisplayMode('table')
  + panel.timeSeries.options.legend.withPlacement('bottom')
  + panel.timeSeries.options.legend.withCalcs(['mean', 'lastNotNull'])
  + panel.timeSeries.options.legend.withSortBy('Mean')
  + panel.timeSeries.options.legend.withSortDesc();

{
  core: {
    statRow(y, h): [
      panel.stat.new('Book count')
      + panel.stat.queryOptions.withTargets([
        prometheus.new(
          config.datasource.metrics.uid,
          'sum(hgraber_next_server_book_total{%s}) by (type)' % config.label.filter.service,
        )
        + prometheus.withLegendFormat('{{type}}')
        + prometheus.withInstant(),
      ])
      + panel.stat.standardOptions.withUnit('short')
      + panel.stat.standardOptions.thresholds.withSteps([greenSteps()])
      + panel.stat.gridPos.withH(h)
      + panel.stat.gridPos.withW(12)
      + panel.stat.gridPos.withX(0)
      + panel.stat.gridPos.withY(y)
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
      + panel.stat.gridPos.withH(h)
      + panel.stat.gridPos.withW(12)
      + panel.stat.gridPos.withX(12)
      + panel.stat.gridPos.withY(y)
      + panel.stat.queryOptions.withDatasource(
        config.datasource.metrics.type,
        config.datasource.metrics.uid,
      ),
    ],
    deltaRow(y, h): [
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
      + panel.barChart.gridPos.withH(h)
      + panel.barChart.gridPos.withW(6)
      + panel.barChart.gridPos.withX(0)
      + panel.barChart.gridPos.withY(y)
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
      + panel.barChart.gridPos.withH(h)
      + panel.barChart.gridPos.withW(6)
      + panel.barChart.gridPos.withX(6)
      + panel.barChart.gridPos.withY(y)
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
      + panel.barChart.gridPos.withH(h)
      + panel.barChart.gridPos.withW(6)
      + panel.barChart.gridPos.withX(12)
      + panel.barChart.gridPos.withY(y)
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
      + panel.barChart.gridPos.withH(h)
      + panel.barChart.gridPos.withW(6)
      + panel.barChart.gridPos.withX(18)
      + panel.barChart.gridPos.withY(y)
      + panel.barChart.queryOptions.withDatasource(
        config.datasource.metrics.type,
        config.datasource.metrics.uid,
      ),
    ],
    fsRow(y, h): [
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
      + simpleTSLegend()
      + panel.timeSeries.gridPos.withH(h)
      + panel.timeSeries.gridPos.withW(12)
      + panel.timeSeries.gridPos.withX(0)
      + panel.timeSeries.gridPos.withY(y)
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
      + simpleTSLegend()
      + panel.timeSeries.gridPos.withH(h)
      + panel.timeSeries.gridPos.withW(12)
      + panel.timeSeries.gridPos.withX(12)
      + panel.timeSeries.gridPos.withY(y)
      + panel.timeSeries.queryOptions.withDatasource(
        config.datasource.metrics.type,
        config.datasource.metrics.uid,
      ),
    ],
  },
  booksAndPages(y):
    [
      panel.row.new('Books and pages')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels([
        panel.timeSeries.new('Book delta count')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(delta(hgraber_next_server_book_total{%s}[$__rate_interval])) by (type)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(0)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Page delta count')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(delta(hgraber_next_server_page_total{%s}[$__rate_interval])) by (type)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(12)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('File delta count')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(delta(hgraber_next_server_file_total{%s}[$__rate_interval])) by (type, fs_id)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}} -> {{fs_id}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(0)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('File delta size')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(delta(hgraber_next_server_file_bytes{%s}[$__rate_interval])) by (type, fs_id)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}} -> {{fs_id}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('bytes')
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(12)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Book count')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_book_total{%s}) by (type)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(0)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Page count')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_page_total{%s}) by (type)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(12)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('File count')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_file_total{%s}) by (type, fs_id)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}} -> {{fs_id}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(0)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('File size')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_file_bytes{%s}) by (type, fs_id)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{type}} -> {{fs_id}}'),
        ])
        + simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('bytes')
        + panel.timeSeries.gridPos.withW(12)
        + panel.timeSeries.gridPos.withH(9)
        + panel.timeSeries.gridPos.withX(12)
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),

      ]),
    ],
  logs(y, h):
    [
      panel.row.new('Logs')
      + panel.row.gridPos.withY(y),
      panel.logs.new('Logs')
      + panel.logs.queryOptions.withTargets([
        query.loki.new(
          config.datasource.logs.uid,
          '{%s}' % config.label.filter.service,
        ),
      ])
      + panel.logs.options.withSortOrder('Descending')
      + panel.logs.gridPos.withH(h - 1)
      + panel.logs.gridPos.withW(24)
      + panel.logs.gridPos.withX(0)
      + panel.logs.gridPos.withY(y + 1)
      + panel.logs.queryOptions.withDatasource(
        config.datasource.logs.type,
        config.datasource.logs.uid,
      ),
    ],
  panels:
    [
      panel.row.new('Simple info'),
    ]
    + self.core.statRow(1, 3)
    + self.core.deltaRow(4, 9)
    + self.core.fsRow(13, 9)
    + self.logs(22, 10)
    + self.booksAndPages(32),
}

local config = import 'config.libsonnet';
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';
local template = import 'template.libsonnet';

local panel = grafonnet.panel;
local query = grafonnet.query;
local prometheus = grafonnet.query.prometheus;

{
  core: {
    statRow(y, h):
      grafonnet.util.grid.wrapPanels([
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
        + panel.stat.standardOptions.thresholds.withSteps([template.greenSteps()])
        + panel.stat.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.stat.new('Page count')
        + panel.stat.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_page_total{%s}) by (type)' % config.label.filter.service,
          )
          + prometheus.withLegendFormat('{{type}}')
          + prometheus.withInstant(),
        ])
        + panel.stat.standardOptions.withUnit('short')
        + panel.stat.standardOptions.thresholds.withSteps([template.greenSteps()])
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
        + panel.stat.standardOptions.thresholds.withSteps([template.greenSteps()])
        + panel.stat.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.stat.new('FS Compression')
        + panel.stat.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            |||
              1 - 
              sum(hgraber_next_server_file_bytes{type="fs", %s}) by (fs_id)
              /
              sum(hgraber_next_server_file_bytes{type="page", %s}) by (fs_id)
            ||| % [
              config.label.filter.service,
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{fs_id}}')
          + prometheus.withInstant(),
        ])
        + panel.stat.standardOptions.withUnit('percentunit')
        + panel.stat.standardOptions.thresholds.withSteps([template.greenSteps()])
        + panel.stat.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], 6, h, y),
    deltaRow(y, h):
      grafonnet.util.grid.wrapPanels([
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
        + panel.barChart.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], 6, h, y),
    fsRow(y, h):
      grafonnet.util.grid.wrapPanels([
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], 12, h, y),
  },
  booksAndPages(y):
    [
      panel.row.new('Books and pages')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels(grafonnet.util.grid.wrapPanels([
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('bytes')
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
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
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('bytes')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], 12, 9, y + 1)),
    ],
  statistic(y):
    local simpleHM(title, metric, yUnit) =
      panel.heatmap.new(title)
      + panel.heatmap.queryOptions.withTargets([
        prometheus.new(
          config.datasource.metrics.uid,
          'sum(%s{%s}) by (le)' % [
            metric,
            config.label.filter.service,
          ],
        )
        + prometheus.withFormat('heatmap')
        + prometheus.withInterval('$%s' % config.variable.delta.name)
        + prometheus.withLegendFormat('__auto'),
      ])
      + panel.heatmap.options.cellValues.withUnit('short')
      + template.heatmapDefaultColor()
      + panel.heatmap.options.filterValues.withLe(1)
      + panel.heatmap.options.yAxis.withUnit(yUnit)
      + panel.heatmap.queryOptions.withDatasource(
        config.datasource.metrics.type,
        config.datasource.metrics.uid,
      );
    [
      panel.row.new('Statistic')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels(grafonnet.util.grid.wrapPanels([
        simpleHM(
          'Book sizes',
          'hgraber_next_server_statistic_book_size_bucket',
          'bytes',
        ),
        simpleHM(
          'Page sizes',
          'hgraber_next_server_statistic_page_size_bucket',
          'bytes',
        ),
        simpleHM(
          'Page in book',
          'hgraber_next_server_statistic_page_in_book_bucket',
          'short',
        ),
        simpleHM(
          'Page sizes by authors',
          'hgraber_next_server_statistic_pages_size_by_author_bucket',
          'bytes',
        ),
        simpleHM(
          'Book by authors',
          'hgraber_next_server_statistic_books_by_author_bucket',
          'short',
        ),
        simpleHM(
          'Page by authors',
          'hgraber_next_server_statistic_pages_by_author_bucket',
          'short',
        ),
      ], 8, 9, y + 1)),
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
  workersAndAgents(y):
    [
      panel.row.new('Workers and agents')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels(grafonnet.util.grid.wrapPanels([
        panel.timeSeries.new('Agent parsing avg latency')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            |||
              sum(rate(hgraber_next_agent_parser_action_seconds_sum{%s}[$__rate_interval])) by (action, parser_name)
              /
              sum(rate(hgraber_next_agent_parser_action_seconds_count{%s}[$__rate_interval])) by (action, parser_name)
            ||| % [
              config.label.filter.service,
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{parser_name}} -> {{action}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('s')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Agent parsing RPS')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(rate(hgraber_next_agent_parser_action_seconds_count{%s}[$__rate_interval])) by (action, parser_name)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{parser_name}} -> {{action}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('reqps')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Worker avg latency')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            |||
              sum(rate(hgraber_next_server_worker_execution_task_seconds_sum{%s}[$__rate_interval])) by (worker_name)
              /
              sum(rate(hgraber_next_server_worker_execution_task_seconds_count{%s}[$__rate_interval])) by (worker_name)
            ||| % [
              config.label.filter.service,
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{worker_name}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('s')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Worker RPS')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(rate(hgraber_next_server_worker_execution_task_seconds_count{%s}[$__rate_interval])) by (worker_name)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{worker_name}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('reqps')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Workers stat')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_worker_total{%s}) by (worker_name, counter)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{worker_name}} -> {{counter}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Web cache RPS')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(rate(hgraber_next_agent_web_cache_total{%s}[$__rate_interval])) by (action)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{action}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('reqps')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], 12, 9, y + 1)),
    ],
  other(y):
    [
      panel.row.new('Other')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels(grafonnet.util.grid.wrapPanels([
        panel.timeSeries.new('FS Compression')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            |||
              1 - 
              sum(hgraber_next_server_file_bytes{type="fs", %s}) by (fs_id)
              /
              sum(hgraber_next_server_file_bytes{type="page", %s}) by (fs_id)
            ||| % [
              config.label.filter.service,
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{fs_id}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('percentunit')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
        panel.timeSeries.new('Metric info collect time')
        + panel.timeSeries.queryOptions.withTargets([
          prometheus.new(
            config.datasource.metrics.uid,
            'sum(hgraber_next_server_info_scrape_duration_seconds{%s}) by (scrape_name)' % [
              config.label.filter.service,
            ],
          )
          + prometheus.withLegendFormat('{{scrape_name}}'),
        ])
        + template.simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('s')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], 12, 9, y + 1)),
    ],
  httpServer(y):
    [
      panel.row.new('HTTP server')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels(grafonnet.util.grid.wrapPanels([
        template.timeSeries(
          'Latency Q95',
          'histogram_quantile(0.95, sum(rate(hgraber_next_http_server_handle_duration_bucket{%s}[$__rate_interval])) by (server_addr, operation, status, le))' % [
            config.label.filter.service,
          ],
          's',
          '{{server_addr}}/{{operation}} -> {{status}}',
        ),
        template.timeSeries(
          'RPS',
          'sum(irate(hgraber_next_http_server_handle_duration_count{%s}[$__rate_interval])) by (server_addr, operation, status)' % [
            config.label.filter.service,
          ],
          'reqps',
          '{{server_addr}}/{{operation}} -> {{status}}',
        ),
        template.timeSeries(
          'Active request',
          'sum(hgraber_next_http_server_active_handlers_total{%s}) by (server_addr, operation)' % [
            config.label.filter.service,
          ],
          'short',
          '{{server_addr}}/{{operation}}',
        ),
        template.table(
          'Slow request',
          'sum(hgraber_next_http_server_handle_duration_sum{%s} / hgraber_next_http_server_handle_duration_count{%s}) by (server_addr, operation, status)' % [
            config.label.filter.service,
            config.label.filter.service,
          ],
          's',
          '{{server_addr}}/{{operation}} -> {{status}}',
        )
        + template.tableTimeOverride(),
      ], 12, 9, y + 1)),
    ],
  database(y):
    [
      panel.row.new('Database')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels(grafonnet.util.grid.wrapPanels([
        template.timeSeries(
          'Latency Q95',
          'histogram_quantile(0.95, sum(rate(hgraber_next_database_request_duration_bucket{%s}[$__rate_interval])) by (stmt, le))' % [
            config.label.filter.service,
          ],
          's',
          '{{stmt}}',
        ),
        template.timeSeries(
          'RPS',
          'sum(irate(hgraber_next_database_request_duration_count{%s}[$__rate_interval])) by (stmt)' % [
            config.label.filter.service,
          ],
          'reqps',
          '{{stmt}}',
        ),
        template.timeSeries(
          'Active request',
          'sum(hgraber_next_database_active_request{%s})' % [
            config.label.filter.service,
          ],
          'short',
        ),
        template.timeSeries(
          'Open connection',
          'sum(hgraber_next_database_open_connection{%s})' % [
            config.label.filter.service,
          ],
          'short',
        ),
        template.table(
          'Slow request',
          'sum(hgraber_next_database_request_duration_sum{%s} / hgraber_next_database_request_duration_count{%s}) by (stmt)' % [
            config.label.filter.service,
            config.label.filter.service,
          ],
          's',
          '{{stmt}}',
        )
        + template.tableTimeOverride(),
      ], 12, 9, y + 1)),
    ],
  panels:
    [
      panel.row.new('Simple info'),
    ]
    + self.core.statRow(1, 3)
    + self.core.deltaRow(4, 6)
    + self.core.fsRow(13, 9)
    + self.logs(22, 10)
    + self.booksAndPages(32)
    + self.statistic(33)
    + self.workersAndAgents(34)
    + self.other(35)
    + self.httpServer(36)
    + self.database(37),
}

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
local heatmapDefaultColor() =
  panel.heatmap.options.color.withFill('semi-dark-blue')
  + panel.heatmap.options.color.withScheme('Blues')
  + panel.heatmap.options.color.withMode('scheme');

{
  utils: {
    autoWX(arr, count=2):
      local mutator(i, e) =
        e
        + panel.timeSeries.gridPos.withW(24 / count)
        + panel.timeSeries.gridPos.withX(24 * (i % count) / count);

      std.mapWithIndex(mutator, arr),
    setH(arr, h):
      local mutator(e) = e + panel.timeSeries.gridPos.withH(h);
      std.map(mutator, arr),
    setY(arr, y):
      local mutator(e) = e + panel.timeSeries.gridPos.withY(y);
      std.map(mutator, arr),
    makeRow(arr, h=9, y=0, onRow=std.length(arr)):
      self.setY(
        self.setH(
          self.autoWX(arr, onRow),
          h,
        ),
        y,
      ),
    makeBlock(arr, h=9, onRow=2):
      self.setH(
        self.autoWX(arr, onRow),
        h,
      ),
  },
  core: {
    statRow(y, h):
      $.utils.makeRow([
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
        + panel.stat.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], h, y),
    deltaRow(y, h):
      $.utils.makeRow([
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
      ], h, y),
    fsRow(y, h):
      $.utils.makeRow([
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
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ], h, y),
  },
  booksAndPages(y):
    [
      panel.row.new('Books and pages')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels($.utils.makeBlock([
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
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ])),
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
      + heatmapDefaultColor()
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
      + panel.row.withPanels($.utils.makeBlock([
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
      ], onRow=3)),
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
      + panel.row.withPanels($.utils.makeBlock([
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
        + simpleTSLegend()
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
        + simpleTSLegend()
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
        + simpleTSLegend()
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
        + simpleTSLegend()
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
        + simpleTSLegend()
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ])),
    ],
  other(y):
    [
      panel.row.new('Other')
      + panel.row.gridPos.withY(y)
      + panel.row.withCollapsed()
      + panel.row.withPanels($.utils.makeBlock([
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
        + simpleTSLegend()
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
        + simpleTSLegend()
        + panel.timeSeries.standardOptions.withUnit('s')
        + panel.timeSeries.queryOptions.withDatasource(
          config.datasource.metrics.type,
          config.datasource.metrics.uid,
        ),
      ])),
    ],
  panels:
    [
      panel.row.new('Simple info'),
    ]
    + self.core.statRow(1, 3)
    + self.core.deltaRow(4, 9)
    + self.core.fsRow(13, 9)
    + self.logs(22, 10)
    + self.booksAndPages(32)
    + self.statistic(33)
    + self.workersAndAgents(34)
    + self.other(35),
}

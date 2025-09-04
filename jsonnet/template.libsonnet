local config = import 'config.libsonnet';
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

local panel = grafonnet.panel;
local query = grafonnet.query;
local prometheus = grafonnet.query.prometheus;


{
  greenSteps():
    { color: 'green', value: null },
  simpleTSLegend():
    panel.timeSeries.options.legend.withDisplayMode('table')
    + panel.timeSeries.options.legend.withPlacement('bottom')
    + panel.timeSeries.options.legend.withCalcs(['mean', 'lastNotNull'])
    + panel.timeSeries.options.legend.withSortBy('Mean')
    + panel.timeSeries.options.legend.withSortDesc(),
  heatmapDefaultColor():
    panel.heatmap.options.color.withFill('semi-dark-blue')
    + panel.heatmap.options.color.withScheme('Blues')
    + panel.heatmap.options.color.withMode('scheme'),
  timeSeries(title, request, unit, legend='__auto'):
    panel.timeSeries.new(title)
    + panel.timeSeries.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        request,
      )
      + prometheus.withLegendFormat(legend),
    ])
    + panel.timeSeries.standardOptions.withUnit(unit)
    + self.simpleTSLegend()
    + panel.timeSeries.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
  table(title, request, unit, legend='__auto'):
    panel.table.new(title)
    + panel.table.queryOptions.withTargets([
      prometheus.new(
        config.datasource.metrics.uid,
        request,
      )
      + prometheus.withLegendFormat(legend)
      + prometheus.withFormat('table')
      + prometheus.withInstant(true),
    ])
    + panel.table.standardOptions.withUnit(unit)
    + panel.table.options.withSortBy([
      panel.table.options.sortBy.withDisplayName('Value')
      + panel.table.options.sortBy.withDesc(true),
    ])
    + panel.table.queryOptions.withDatasource(
      config.datasource.metrics.type,
      config.datasource.metrics.uid,
    ),
  tableTimeOverride():
    panel.table.standardOptions.withOverrides([
      panel.table.standardOptions.override.byName.new('Time')
      + panel.table.standardOptions.override.byName.withProperty('custom.hidden', true),
    ]),
}

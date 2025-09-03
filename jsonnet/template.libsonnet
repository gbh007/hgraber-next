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
}

local config = import 'config.libsonnet';
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';
local panel = import 'panel.libsonnet';
local variable = import 'variable.libsonnet';

grafonnet.dashboard.new(config.dashboard.title)
+ grafonnet.dashboard.withUid(config.dashboard.uid)
+ grafonnet.dashboard.withTimezone(config.dashboard.timezone)
+ grafonnet.dashboard.withWeekStart(config.dashboard.weekStart)
+ grafonnet.dashboard.time.withFrom(config.dashboard.timeFrom)
+ grafonnet.dashboard.withRefresh(config.dashboard.refresh)
+ grafonnet.dashboard.graphTooltip.withSharedCrosshair()
+ grafonnet.dashboard.withTags(config.dashboard.tags)
+ grafonnet.dashboard.withLinks([
  grafonnet.dashboard.link.link.new('GitHub', 'https://github.com/gbh007/hgraber-next')
  + grafonnet.dashboard.link.link.options.withTargetBlank(),
  grafonnet.dashboard.link.dashboards.new('HG next boards', config.dashboard.tags)
  + grafonnet.dashboard.link.dashboards.options.withAsDropdown()
  + grafonnet.dashboard.link.dashboards.options.withKeepTime(),
])
+ grafonnet.dashboard.withAnnotations([
  {
    datasource: {
      type: config.datasource.logs.type,
      uid: config.datasource.logs.uid,
    },
    enable: false,
    expr: '{%s} |= `application start`' % config.label.filter.service,
    iconColor: 'super-light-purple',
    name: 'app started (logs)',
    tagKeys: '{{host}}',
    textFormat: 'started',
    titleFormat: '{{%s}}' % config.variable.serviceName.name,
  },
  {
    datasource: {
      type: config.datasource.metrics.type,
      uid: config.datasource.metrics.uid,
    },
    enable: true,
    expr: 'hgraber_next_agent_version_info{%s} * 1000 or hgraber_next_server_version_info{%s} * 1000' % [
      config.label.filter.service,
      config.label.filter.service,
    ],
    iconColor: 'super-light-blue',
    name: 'app started (metrics)',
    tagKeys: '{{host}}',
    textFormat: 'started',
    titleFormat: '{{%s}}' % config.variable.serviceName.name,
    useValueForTime: 'on',
  },
  // FIXME: генерировать полностью с помощью либы.
  // + grafonnet.dashboard.annotation.list.withName('app started')
  // + grafonnet.dashboard.annotation.list.withEnable()
  // + grafonnet.dashboard.annotation.list.datasource.withType(config.datasource.logs.type)
  // + grafonnet.dashboard.annotation.list.datasource.withUid(config.datasource.logs.uid)
  // + grafonnet.dashboard.annotation.list.withExpr('{%s} |= `application start`' % config.label.filter.service)
  // + grafonnet.dashboard.annotation.list.withIconColor('super-light-blue'),
])
+ grafonnet.dashboard.withPanels(panel.panels)
+ grafonnet.dashboard.withVariables([
  variable.logs(),
  variable.metrics(),
  variable.service(),
  variable.delta(),
])

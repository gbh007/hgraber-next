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
+ grafonnet.dashboard.withPanels(panel.panels)
+ grafonnet.dashboard.withVariables([
  variable.logs(),
  variable.metrics(),
  variable.service(),
  variable.delta(),
])

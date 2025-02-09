{
  dashboard: {
    title: 'Hgraber next',
    uid: 'fecigwx9lk1z4d',  // KpCSmwlIk
    timezone: 'Asia/Krasnoyarsk',
    weekStart: 'monday',
    timeFrom: 'now-6h',
    refresh: '1m',
    tags: ['hgnext'],
  },
  variable: {
    prometheus: {
      type: 'prometheus',
      name: 'metrics',
    },
    loki: {
      type: 'loki',
      name: 'logs',
    },
    serviceName: {
      name: 'service_name',
      current: if std.extVar('services') != '' then std.split(std.extVar('services'), ',') else [],
    },
    delta: {
      name: 'deltaInterval',
      values: ['1m', '5m', '10m', '30m', '1h', '4h', '8h', '1d', '7d'],
      current: '4h',
    },
  },
  label: {
    filter: {
      service: 'service_name=~"$%s"' % $.variable.serviceName.name,
    },
  },
  datasource: {
    metrics: {
      type: $.variable.prometheus.type,
      uid: '${%s}' % $.variable.prometheus.name,
    },
    logs: {
      type: $.variable.loki.type,
      uid: '${%s}' % $.variable.loki.name,
    },
  },
}

package metricagent

import (
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ParserActionTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: ParserActionSecondsName,
		Help: "Время действий парсинга",
	}, []string{metriccore.ActionLabel, ParserNameLabel})

	WebCacheCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: WebCacheTotalName,
		Help: "Количество действий с кешом для реквестера",
	}, []string{metriccore.ActionLabel})
)

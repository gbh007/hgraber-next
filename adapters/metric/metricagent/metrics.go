package metricagent

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
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

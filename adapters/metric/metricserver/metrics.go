package metricserver

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
)

var (
	BookTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: BookTotalName,
		Help: "Количество книг по статусам",
	}, []string{metriccore.TypeLabel})
	PageTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: PageTotalName,
		Help: "Количество страниц по статусам",
	}, []string{metriccore.TypeLabel})
	LastCollectorScrapeDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: LastCollectorScrapeDurationName,
		Help: "Время последней сборки данных",
	}, []string{ScrapeNameLabel})
)

package metrichttp

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ServerHandleRequest = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    ServerHandleDurationName,
			Buckets: []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.2, 0.5, 1, 1.5, 2, 3, 5, 10},
			Help:    "Время обработки запроса",
		},
		[]string{ServerAddrLabelName, OperationLabelName, StatusLabelName},
	)
	ServerActiveRequest = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: ServerActiveHandlersTotalName,
			Help: "Количество активных запросов",
		},
		[]string{ServerAddrLabelName, OperationLabelName},
	)
)

package metricfs

import (
	"github.com/gbh007/hgraber-next/adapters/metric/metriccore"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ActionTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: ActionSecondsName,
		Help: "Время действий с файловой системой",
	}, []string{metriccore.ActionLabel, FSIDLabel})

	FileTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: FileTotalName,
		Help: "Количество файлов по статусам",
	}, []string{metriccore.TypeLabel, FSIDLabel})

	FileBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: FileBytesName,
		Help: "Размер файлов по статусам",
	}, []string{metriccore.TypeLabel, FSIDLabel})
)

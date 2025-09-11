package metricserver

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/metrics/metriccore"
)

var (
	BookTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"book_total",
	)
	PageTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"page_total",
	)
	FileTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"file_total",
	)
	FileBytesName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"file_bytes",
	)
	LastCollectorScrapeDurationName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"info_scrape_duration_seconds",
	)
	FSActionSecondsName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"fs_action_seconds",
	)
	WorkerExecutionTaskSecondsName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"worker_execution_task_seconds",
	)
	WorkerTotalName = prometheus.BuildFQName(
		metriccore.SystemName,
		metriccore.SubSystemName,
		"worker_total",
	)
)

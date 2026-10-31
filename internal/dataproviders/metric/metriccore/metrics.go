package metriccore

import (
	"github.com/gbh007/hgraber-next/internal/version"
	"github.com/prometheus/client_golang/prometheus"
)

var VersionInfoMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: VersionInfoName,
	Help: "Информация о приложении",
	ConstLabels: prometheus.Labels{
		"go_version": version.GoVersion,
		"go_os":      version.GoOS,
		"go_arch":    version.GoArch,
		"version":    version.Version,
		"commit":     version.Commit,
		"branch":     version.Branch,
		"build":      version.BuildAt,
	},
})

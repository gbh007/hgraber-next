package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/gbh007/hgraber-next/internal/version"
)

var versionInfo = promauto.NewGauge(prometheus.GaugeOpts{
	Namespace: SystemName,
	Subsystem: SubSystemName,
	Name:      "version_info",
	Help:      "Информация о приложении",
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

func init() {
	versionInfo.Set(1)
}

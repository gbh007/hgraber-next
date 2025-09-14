package metric

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/gbh007/hgraber-next/adapters/metric/metricserver"
	"github.com/gbh007/hgraber-next/version"
)

var versionInfo = promauto.NewGauge(prometheus.GaugeOpts{
	Name: metricserver.VersionInfoName,
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

func init() {
	versionInfo.Set(float64(time.Now().Unix()))
}

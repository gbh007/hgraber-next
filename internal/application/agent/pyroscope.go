package agent

import (
	"log/slog"
	"runtime"

	"github.com/grafana/pyroscope-go"

	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func initPyroscope[T any](logger *slog.Logger, cfg agnconfig.AgentConfig[T]) (*pyroscope.Profiler, error) {
	runtime.SetMutexProfileFraction(cfg.Application.Pyroscope.Rate)
	runtime.SetBlockProfileRate(cfg.Application.Pyroscope.Rate)

	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: cfg.Application.ServiceName,
		ServerAddress:   cfg.Application.Pyroscope.Endpoint,

		Logger: pkg.NewPyroscopeLogger(logger, cfg.Application.Pyroscope.Debug),

		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		return nil, err
	}

	return profiler, nil
}

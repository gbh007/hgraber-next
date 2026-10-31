package server

import (
	"fmt"
	"log/slog"
	"runtime"

	"github.com/gbh007/hgraber-next/internal/config"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/grafana/pyroscope-go"
)

func initPyroscope(logger *slog.Logger, cfg config.ServerConfig) (*pyroscope.Profiler, error) {
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
		return nil, fmt.Errorf("start pyroscope: %w", err)
	}

	return profiler, nil
}

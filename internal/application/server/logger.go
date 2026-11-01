package server

import (
	"log/slog"
	"os"

	"github.com/gbh007/hgraber-next/internal/application/server/srvconfig"
)

func initLogger(cfg srvconfig.ServerConfig) *slog.Logger {
	slogOpt := &slog.HandlerOptions{
		AddSource: cfg.Log.IncludeSource,
		Level:     cfg.Log.SlogLevel(),
	}

	logger := slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			slogOpt,
		),
	)

	if cfg.Application.ServiceName != "" {
		logger = logger.With(slog.String("service_name", cfg.Application.ServiceName))
	}

	return logger
}

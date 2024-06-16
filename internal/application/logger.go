package application

import (
	"log/slog"
	"os"
)

func initLogger(cfg Config) *slog.Logger {
	slogOpt := &slog.HandlerOptions{
		AddSource: cfg.Debug,
		Level:     slog.LevelInfo,
	}

	if cfg.Debug {
		slogOpt.Level = slog.LevelDebug
	}

	return slog.New(slog.NewJSONHandler(
		os.Stderr,
		slogOpt,
	))
}

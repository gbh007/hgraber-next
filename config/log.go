package config

import (
	"log/slog"
	"strings"
)

type Log struct {
	IncludeSource bool   `toml:"include_source" yaml:"include_source" envconfig:"INCLUDE_SOURCE"`
	Level         string `toml:"level" yaml:"level" envconfig:"LEVEL"`
}

func LogDefault() Log {
	return Log{
		Level: "info",
	}
}

func (l Log) SlogLevel() slog.Level {
	switch strings.ToLower(l.Level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

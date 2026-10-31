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
	case "debug", "dbg":
		return slog.LevelDebug
	case "info", "inf":
		return slog.LevelInfo
	case "warn", "warning", "wrn":
		return slog.LevelWarn
	case "error", "err":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

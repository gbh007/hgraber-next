package agnconfig

import (
	"log/slog"
	"strings"
)

type AgentLog struct {
	IncludeSource bool   `toml:"include_source" yaml:"include_source" envconfig:"INCLUDE_SOURCE"`
	Level         string `toml:"level" yaml:"level" envconfig:"LEVEL"`
	RequestErrors bool   `toml:"request_errors" yaml:"request_errors" envconfig:"REQUEST_ERRORS"`
}

func AgentLogDefault() AgentLog {
	return AgentLog{
		Level: "info",
	}
}

func (l AgentLog) SlogLevel() slog.Level {
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

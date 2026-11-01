package agnconfig

import (
	"log/slog"
	"strings"
)

// FIXME: слить с серверным
type AgentLog struct {
	IncludeSource bool   `toml:"include_source"`
	Level         string `toml:"level"`
	RequestErrors bool   `toml:"request_errors"`
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

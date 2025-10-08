package pkg

import (
	"fmt"
	"log/slog"

	"github.com/grafana/pyroscope-go"
)

var _ pyroscope.Logger = (*PyroscopeLogger)(nil)

type PyroscopeLogger struct {
	logger *slog.Logger
	debug  bool
}

func NewPyroscopeLogger(logger *slog.Logger, debug bool) *PyroscopeLogger {
	return &PyroscopeLogger{
		logger: logger,
		debug:  debug,
	}
}

func (l PyroscopeLogger) Infof(msg string, args ...any) {
	l.logger.Info(fmt.Sprintf(msg, args...)) //nolint:sloglint // особенность реализации
}

func (l PyroscopeLogger) Debugf(msg string, args ...any) {
	if !l.debug {
		return
	}

	l.logger.Debug(fmt.Sprintf(msg, args...)) //nolint:sloglint // особенность реализации
}

func (l PyroscopeLogger) Errorf(msg string, args ...any) {
	l.logger.Error(fmt.Sprintf(msg, args...)) //nolint:sloglint // особенность реализации
}

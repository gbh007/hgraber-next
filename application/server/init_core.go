package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel"

	"github.com/gbh007/hgraber-next/adapters/metric"
	"github.com/gbh007/hgraber-next/controllers/async"
	"github.com/gbh007/hgraber-next/controllers/workermanager"
)

func (a *App) initLogger() {
	a.Logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
}

func (a *App) initCore(ctx context.Context) error {
	var err error

	// Инициализируем логгер по умолчанию
	a.initLogger()

	a.Config, err = parseConfig()
	if err != nil {
		return fmt.Errorf("fail parse config: %w", err)
	}

	// Переинициализируем логгер по параметрам конфигурации
	a.Logger = initLogger(a.Config)

	a.metricProvider, err = metric.New(metric.Config{
		ServiceName:    a.Config.Application.ServiceName,
		Type:           metric.ServerSystemType,
		WithGo:         true,
		WithVersion:    true,
		WithFS:         true,
		WithServer:     true,
		WithDB:         true,
		WithHTTPServer: true,
		WithAgent:      false,
	})
	if err != nil {
		return fmt.Errorf("fail init metrics: %w", err)
	}

	if a.Config.Application.Pyroscope.Endpoint != "" {
		a.pyroscopeProfiler, err = initPyroscope(a.Logger, a.Config)
		if err != nil {
			return fmt.Errorf("fail init pyroscope: %w", err)
		}

		defer a.pyroscopeProfiler.Stop() //nolint:errcheck // будет исправлено позднее
	}

	if a.Config.Application.TraceEndpoint != "" {
		err := initTrace(ctx, a.Config.Application.TraceEndpoint, a.Config.Application.ServiceName)
		if err != nil {
			return fmt.Errorf("fail init otel: %w", err)
		}
	}

	a.Tracer = otel.GetTracerProvider().Tracer("hgraber-next")

	a.asyncController = async.New(a.Logger)

	a.workersController = workermanager.New(a.Logger)
	a.asyncController.RegisterRunner(a.workersController)

	return nil
}

//revive:disable:file-length-limit
package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/grafana/pyroscope-go"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"

	"github.com/gbh007/hgraber-next/adapters/agent"
	"github.com/gbh007/hgraber-next/adapters/filestorage"
	"github.com/gbh007/hgraber-next/adapters/metric"
	"github.com/gbh007/hgraber-next/adapters/postgresql"
	"github.com/gbh007/hgraber-next/adapters/tmpdata"
	"github.com/gbh007/hgraber-next/config"
	"github.com/gbh007/hgraber-next/controllers/async"
	"github.com/gbh007/hgraber-next/controllers/workermanager"
	"github.com/gbh007/hgraber-next/usecases/agentusecase"
	"github.com/gbh007/hgraber-next/usecases/attributeusecase"
	"github.com/gbh007/hgraber-next/usecases/bffusecase"
	"github.com/gbh007/hgraber-next/usecases/bookusecase"
	"github.com/gbh007/hgraber-next/usecases/cleanupusecase"
	"github.com/gbh007/hgraber-next/usecases/deduplicatorusecase"
	"github.com/gbh007/hgraber-next/usecases/exportusecase"
	"github.com/gbh007/hgraber-next/usecases/filesystemusecase"
	"github.com/gbh007/hgraber-next/usecases/hproxyusecase"
	"github.com/gbh007/hgraber-next/usecases/labelusecase"
	"github.com/gbh007/hgraber-next/usecases/massloadusecase"
	"github.com/gbh007/hgraber-next/usecases/parsingusecase"
	"github.com/gbh007/hgraber-next/usecases/rebuilderusecase"
	"github.com/gbh007/hgraber-next/usecases/systemusecase"
)

type App struct {
	Config config.Config
	Logger *slog.Logger
	Tracer trace.Tracer

	pyroscopeProfiler *pyroscope.Profiler

	metricProvider     *metric.MetricProvider
	storage            *postgresql.Repo
	agentSystem        *agent.Client
	fileStorageAdapter *filestorage.Storage
	tmpStorage         *tmpdata.Storage

	bookUseCases        *bookusecase.UseCase
	parsingUseCases     *parsingusecase.UseCase
	exportUseCases      *exportusecase.UseCase
	deduplicateUseCases *deduplicatorusecase.UseCase
	cleanupUseCases     *cleanupusecase.UseCase
	reBuilderUseCases   *rebuilderusecase.UseCase
	fsUseCases          *filesystemusecase.UseCase
	bffUseCases         *bffusecase.UseCase
	attributeUseCases   *attributeusecase.UseCase
	labelUseCases       *labelusecase.UseCase
	massloadUseCases    *massloadusecase.UseCase
	systemUseCases      *systemusecase.UseCase
	agentUseCases       *agentusecase.UseCase
	hProxyUseCases      *hproxyusecase.UseCase

	asyncController   *async.Controller
	workersController *workermanager.Controller
}

func New() *App {
	a := &App{
		Tracer: noop.Tracer{},
	}
	a.initLogger()

	return a
}

func (a *App) Close() error {
	if a.pyroscopeProfiler != nil {
		err := a.pyroscopeProfiler.Stop()
		if err != nil {
			return fmt.Errorf("stop pyroscope: %w", err)
		}
	}

	return nil
}

func (a *App) Init(ctx context.Context) error {
	var err error

	err = a.initCore(ctx)
	if err != nil {
		return fmt.Errorf("fail init core: %w", err)
	}

	err = a.initAdapters(ctx)
	if err != nil {
		return fmt.Errorf("fail init adapters: %w", err)
	}

	a.initUseCases()

	err = a.initControllers(ctx)
	if err != nil {
		return fmt.Errorf("fail init controllers: %w", err)
	}

	return nil
}

func (a *App) Serve(ctx context.Context) error {
	a.Logger.InfoContext(ctx, "application start")
	defer a.Logger.InfoContext(ctx, "application stop")

	err := a.asyncController.Serve(ctx)
	if err != nil {
		return fmt.Errorf("fail run controllers: %w", err)
	}

	return nil
}

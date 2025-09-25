//revive:disable:file-length-limit
package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/gbh007/hgraber-next/adapters/agent"
	"github.com/gbh007/hgraber-next/adapters/filestorage"
	"github.com/gbh007/hgraber-next/adapters/metric"
	"github.com/gbh007/hgraber-next/adapters/postgresql"
	"github.com/gbh007/hgraber-next/adapters/tmpdata"
	"github.com/gbh007/hgraber-next/controllers/apiagent"
	"github.com/gbh007/hgraber-next/controllers/apiserver"
	"github.com/gbh007/hgraber-next/controllers/workermanager"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
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

//nolint:cyclop,funlen // будет исправлено позднее
func Serve() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	cfg, err := parseConfig()
	if err != nil {
		// Поскольку на этот момент нет ни логгера ни вообще ничего то выкидываем панику.
		panic(err)
	}

	logger := initLogger(cfg)

	metricProvider, err := metric.New(metric.Config{
		ServiceName:    cfg.Application.ServiceName,
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
		logger.ErrorContext(
			ctx, "fail init metrics",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	if cfg.Application.Pyroscope.Endpoint != "" {
		profiler, err := initPyroscope(logger, cfg)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init pyroscope",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		defer profiler.Stop() //nolint:errcheck // будет исправлено позднее
	}

	if cfg.Application.TraceEndpoint != "" {
		err := initTrace(ctx, cfg.Application.TraceEndpoint, cfg.Application.ServiceName)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init otel",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	tracer := otel.GetTracerProvider().Tracer("hgraber-next")

	asyncController := New(logger)
	tmpStorage := tmpdata.New()

	storage, err := postgresql.New(
		ctx,
		logger,
		tracer,
		metricProvider,
		cfg.Storage.DebugPGX,
		cfg.Storage.Connection,
		cfg.Storage.MaxConnections,
	)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init postgres",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	agents, err := storage.Agents(ctx, core.AgentFilter{})
	if err != nil {
		logger.ErrorContext(
			ctx, "fail load agents from storage",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	agentSystem, err := agent.New(agents, cfg.Parsing.AgentTimeout)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init agent system",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	fileStorageAdapter := filestorage.New(
		logger,
		agentSystem,
		storage,
		metricProvider,
		cfg.FileStorage.TryReconnect,
	)

	err = fileStorageAdapter.Init(ctx, true)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init file system",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	bookUseCases := bookusecase.New(logger, storage)
	parsingUseCases := parsingusecase.New(
		logger,
		storage,
		agentSystem,
		fileStorageAdapter,
		bookUseCases,
		cfg.Parsing.ParseBookTimeout,
		cfg.AttributeRemap.Auto,
		cfg.AttributeRemap.AllLower,
	)
	exportUseCases := exportusecase.New(
		logger,
		storage,
		fileStorageAdapter,
		agentSystem,
		tmpStorage,
		bookUseCases,
		cfg.AttributeRemap.Auto,
		cfg.AttributeRemap.AllLower,
	)
	deduplicateUseCases := deduplicatorusecase.New(logger, storage, tracer)
	cleanupUseCases := cleanupusecase.New(logger, tracer, storage, fileStorageAdapter)
	reBuilderUseCases := rebuilderusecase.New(
		logger,
		tracer,
		storage,
		cfg.AttributeRemap.Auto,
		cfg.AttributeRemap.AllLower,
	)
	fsUseCases := filesystemusecase.New(logger, storage, fileStorageAdapter, tmpStorage)
	bffUseCases := bffusecase.New(logger, storage, deduplicateUseCases)
	attributeUseCases := attributeusecase.New(logger, storage, cfg.AttributeRemap.AllLower)
	labelUseCases := labelusecase.New(logger, storage)
	massloadUseCases := massloadusecase.New(logger, storage, tmpStorage, agentSystem)

	workerUnits := make([]workermanager.WorkerUnit, 0, len(cfg.Workers))

	for _, wCfg := range cfg.Workers {
		var unit workermanager.WorkerUnit

		switch wCfg.GetName() {
		case systemmodel.WorkerNameBook:
			unit = workermanager.NewBookParser(parsingUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNamePage:
			unit = workermanager.NewPageDownloader(parsingUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameHasher:
			unit = workermanager.NewHasher(fsUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameExporter:
			unit = workermanager.NewExporter(exportUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameTasker:
			unit = workermanager.NewTasker(tmpStorage, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameFileValidator:
			unit = workermanager.NewFileValidator(fsUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameFileTransferer:
			unit = workermanager.NewFileTransfer(fsUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameMassloadSizer:
			unit = workermanager.NewMassloadSize(massloadUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameMassloadAttributeSizer:
			unit = workermanager.NewMassloadAttributeSize(massloadUseCases, logger, tracer, wCfg, metricProvider)

		case systemmodel.WorkerNameMassloadCalculation:
			unit = workermanager.NewMassloadCalculation(massloadUseCases, logger, tracer, wCfg, metricProvider)

		default:
			continue
		}

		workerUnits = append(workerUnits, unit)
	}

	workersController := workermanager.New(logger, workerUnits...)
	asyncController.RegisterRunner(workersController)

	systemUseCases := systemusecase.New(
		logger,
		storage,
		tmpStorage,
		deduplicateUseCases,
		cleanupUseCases,
		workersController,
		attributeUseCases,
	)

	agentUseCases := agentusecase.New(logger, agentSystem, storage)
	hProxyUseCases := hproxyusecase.New(
		logger,
		storage,
		agentSystem,
		cfg.Parsing.ParseBookTimeout,
		cfg.AttributeRemap.Auto,
		cfg.AttributeRemap.AllLower,
	)

	apiController, err := apiserver.New(
		logger,
		tracer,
		cfg.API,
		metricProvider,
		parsingUseCases,
		agentUseCases,
		exportUseCases,
		deduplicateUseCases,
		systemUseCases,
		reBuilderUseCases,
		fsUseCases,
		bffUseCases,
		attributeUseCases,
		labelUseCases,
		bookUseCases,
		hProxyUseCases,
		massloadUseCases,
	)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail to create api server",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	asyncController.RegisterRunner(apiController)

	if cfg.AgentServer.Addr != "" {
		apiAgentController, err := apiagent.New(
			cfg.AgentServer,
			time.Now(),
			logger,
			tracer,
			parsingUseCases,
			exportUseCases,
			metricProvider,
		)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail to create api agent",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		asyncController.RegisterRunner(apiAgentController)
	}

	if cfg.Application.Metric.Enabled() {
		infoCollector, err := metricProvider.NewSystemInfoCollector(
			logger,
			systemUseCases,
			storage,
			cfg.Application.Metric,
		)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail to create info collector",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		asyncController.RegisterRunner(infoCollector)
	}

	logger.InfoContext(ctx, "application start")
	defer logger.InfoContext(ctx, "application stop")

	err = asyncController.Serve(ctx)
	if err != nil {
		logger.ErrorContext(
			ctx, "serve application error",
			slog.Any("error", err),
		)

		os.Exit(1)
	}
}

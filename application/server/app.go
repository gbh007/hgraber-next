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
	"github.com/gbh007/hgraber-next/adapters/fileStorage"
	"github.com/gbh007/hgraber-next/adapters/postgresql"
	"github.com/gbh007/hgraber-next/adapters/tmpdata"
	"github.com/gbh007/hgraber-next/controllers/apiagent"
	"github.com/gbh007/hgraber-next/controllers/apiserver"
	"github.com/gbh007/hgraber-next/controllers/workermanager"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/metrics"
	"github.com/gbh007/hgraber-next/usecases/agentusecase"
	"github.com/gbh007/hgraber-next/usecases/attributeusecase"
	"github.com/gbh007/hgraber-next/usecases/bffusecase"
	"github.com/gbh007/hgraber-next/usecases/bookusecase"
	"github.com/gbh007/hgraber-next/usecases/cleanupusecase"
	"github.com/gbh007/hgraber-next/usecases/deduplicatorusecase"
	"github.com/gbh007/hgraber-next/usecases/exportusecase"
	"github.com/gbh007/hgraber-next/usecases/filesystemusecase"
	"github.com/gbh007/hgraber-next/usecases/labelusecase"
	"github.com/gbh007/hgraber-next/usecases/parsingusecase"
	"github.com/gbh007/hgraber-next/usecases/rebuilderusecase"
	"github.com/gbh007/hgraber-next/usecases/systemusecase"
)

func Serve() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	metricProvider := metrics.MetricProvider{}

	cfg, err := parseConfig()
	if err != nil {
		// Поскольку на этот момент нет ни логгера ни вообще ничего то выкидываем панику.
		panic(err)
	}

	logger := initLogger(cfg)

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
		cfg.Storage.DebugPGX,
		cfg.Storage.DebugSquirrel,
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

	fileStorageAdapter := fileStorage.New(
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
	parsingUseCases := parsingusecase.New(logger, storage, agentSystem, fileStorageAdapter, bookUseCases, cfg.Parsing.ParseBookTimeout, cfg.AttributeRemap.Auto, cfg.AttributeRemap.AllLower)
	exportUseCases := exportusecase.New(logger, storage, fileStorageAdapter, agentSystem, tmpStorage, bookUseCases, cfg.AttributeRemap.Auto, cfg.AttributeRemap.AllLower)
	deduplicateUseCases := deduplicatorusecase.New(logger, storage, tracer)
	cleanupUseCases := cleanupusecase.New(logger, tracer, storage, fileStorageAdapter)
	reBuilderUseCases := rebuilderusecase.New(logger, tracer, storage, cfg.AttributeRemap.Auto, cfg.AttributeRemap.AllLower)
	fsUseCases := filesystemusecase.New(logger, storage, fileStorageAdapter, tmpStorage)
	bffUseCases := bffusecase.New(logger, storage, deduplicateUseCases)
	attributeUseCases := attributeusecase.New(logger, storage, cfg.AttributeRemap.AllLower)
	labelUseCases := labelusecase.New(logger, storage)

	workersController := workermanager.New(
		logger,
		workermanager.NewBookParser(parsingUseCases, logger, tracer, cfg.Workers.Book, metricProvider),
		workermanager.NewPageDownloader(parsingUseCases, logger, tracer, cfg.Workers.Page, metricProvider),
		workermanager.NewHasher(fsUseCases, logger, tracer, cfg.Workers.Hasher, metricProvider),
		workermanager.NewExporter(exportUseCases, logger, tracer, cfg.Workers.Exporter, metricProvider),
		workermanager.NewTasker(tmpStorage, logger, tracer, cfg.Workers.Tasker, metricProvider),
		workermanager.NewFileValidator(fsUseCases, logger, tracer, cfg.Workers.FileValidator, metricProvider),
		workermanager.NewFileTransfer(fsUseCases, logger, tracer, cfg.Workers.FileTransferer, metricProvider),
	)
	asyncController.RegisterRunner(workersController)

	systemUseCases := systemusecase.New(logger, storage, tmpStorage, deduplicateUseCases, cleanupUseCases, workersController, attributeUseCases)

	agentUseCases := agentusecase.New(logger, agentSystem, storage)

	apiController, err := apiserver.New(
		logger,
		tracer,
		cfg.API,
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
		nil, // FIXME
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
		infoCollector, err := metrics.NewSystemInfoCollector(
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

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
	agentUC "github.com/gbh007/hgraber-next/usecases/agent"
	"github.com/gbh007/hgraber-next/usecases/agentcache"
	"github.com/gbh007/hgraber-next/usecases/bff"
	"github.com/gbh007/hgraber-next/usecases/bookrequester"
	"github.com/gbh007/hgraber-next/usecases/cleanup"
	"github.com/gbh007/hgraber-next/usecases/deduplicator"
	"github.com/gbh007/hgraber-next/usecases/export"
	"github.com/gbh007/hgraber-next/usecases/filelogic"
	"github.com/gbh007/hgraber-next/usecases/filesystem"
	"github.com/gbh007/hgraber-next/usecases/parsing"
	"github.com/gbh007/hgraber-next/usecases/rebuilder"
	"github.com/gbh007/hgraber-next/usecases/taskhandler"
	"github.com/gbh007/hgraber-next/usecases/webapi"
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
		cfg.Application.Debug.DB,
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

	err = fileStorageAdapter.InitLegacy(ctx, cfg.Storage.FSAgentID, cfg.Storage.FilePath, false)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init legacy file system",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	err = fileStorageAdapter.Init(ctx, true)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init file system",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	bookRequestUseCases := bookrequester.New(logger, storage)
	parsingUseCases := parsing.New(logger, storage, agentSystem, fileStorageAdapter, bookRequestUseCases, cfg.Parsing.ParseBookTimeout)
	fileUseCases := filelogic.New(logger, storage, fileStorageAdapter)
	exportUseCases := export.New(logger, storage, fileStorageAdapter, agentSystem, tmpStorage, bookRequestUseCases)
	deduplicateUseCases := deduplicator.New(logger, storage, tracer)
	cleanupUseCases := cleanup.New(logger, tracer, storage, fileStorageAdapter)
	taskUseCases := taskhandler.New(logger, tmpStorage, deduplicateUseCases, cleanupUseCases)
	rebuilderUseCases := rebuilder.New(logger, tracer, storage)
	fsUseCases := filesystem.New(logger, storage, fileStorageAdapter, tmpStorage)
	bffUseCases := bff.New(logger, storage)

	workersController := workermanager.New(
		logger,
		workermanager.NewBookParser(parsingUseCases, logger, tracer, cfg.Workers.Book, metricProvider),
		workermanager.NewPageDownloader(parsingUseCases, logger, tracer, cfg.Workers.Page, metricProvider),
		workermanager.NewHasher(fileUseCases, logger, tracer, cfg.Workers.Hasher, metricProvider),
		workermanager.NewExporter(exportUseCases, logger, tracer, cfg.Workers.Exporter, metricProvider),
		workermanager.NewTasker(tmpStorage, logger, tracer, cfg.Workers.Tasker, metricProvider),
		workermanager.NewFileValidator(fsUseCases, logger, tracer, cfg.Workers.FileValidator, metricProvider),
		workermanager.NewFileTransfer(fsUseCases, logger, tracer, cfg.Workers.FileTransferer, metricProvider),
	)
	asyncController.RegisterRunner(workersController)

	webAPIUseCases := webapi.New(
		logger,
		workersController,
		storage,
		fileStorageAdapter,
		bookRequestUseCases,
		deduplicateUseCases,
	)
	agentUseCases := agentUC.New(logger, agentSystem, storage)

	apiController, err := apiserver.New(
		logger,
		tracer,
		cfg.API,
		parsingUseCases,
		webAPIUseCases,
		agentUseCases,
		exportUseCases,
		deduplicateUseCases,
		taskUseCases,
		rebuilderUseCases,
		fsUseCases,
		bffUseCases,
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
		agentCacheUseCase := agentcache.New(logger, parsingUseCases)

		apiAgentController, err := apiagent.New(
			cfg.AgentServer,
			time.Now(),
			logger,
			tracer,
			agentCacheUseCase,
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
			webAPIUseCases,
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

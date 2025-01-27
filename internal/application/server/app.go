package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/otel"

	"hgnext/internal/adapters/agent"
	"hgnext/internal/adapters/fileStorage"
	"hgnext/internal/adapters/postgresql"
	"hgnext/internal/adapters/tmpdata"
	"hgnext/internal/controllers/apiagent"
	"hgnext/internal/controllers/apiserver"
	"hgnext/internal/controllers/workermanager"
	"hgnext/internal/entities"
	"hgnext/internal/metrics"
	agentUC "hgnext/internal/usecases/agent"
	"hgnext/internal/usecases/agentcache"
	"hgnext/internal/usecases/bff"
	"hgnext/internal/usecases/bookrequester"
	"hgnext/internal/usecases/cleanup"
	"hgnext/internal/usecases/deduplicator"
	"hgnext/internal/usecases/export"
	"hgnext/internal/usecases/filelogic"
	"hgnext/internal/usecases/filesystem"
	"hgnext/internal/usecases/parsing"
	"hgnext/internal/usecases/rebuilder"
	"hgnext/internal/usecases/taskhandler"
	"hgnext/internal/usecases/webapi"
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

	agents, err := storage.Agents(ctx, entities.AgentFilter{})
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
		false, // FIXME: вынести в конфиг
	)

	err = fileStorageAdapter.InitLegacy(ctx, cfg.Storage.FSAgentID, cfg.Storage.FilePath, true) // TODO: отключить принудительную часть
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

	metricProvider := metrics.MetricProvider{}

	workersController := workermanager.New(
		logger,
		workermanager.NewBookParser(parsingUseCases, logger, tracer, cfg.Workers.Book, metricProvider),
		workermanager.NewPageDownloader(parsingUseCases, logger, tracer, cfg.Workers.Page, metricProvider),
		workermanager.NewHasher(fileUseCases, logger, tracer, cfg.Workers.Hasher, metricProvider),
		workermanager.NewExporter(exportUseCases, logger, tracer, cfg.Workers.Exporter, metricProvider),
		workermanager.NewTasker(tmpStorage, logger, tracer, cfg.Workers.Tasker, metricProvider),
		workermanager.NewFileValidator(fsUseCases, logger, tracer, cfg.Workers.FileValidator, metricProvider),
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
		cfg.Application.Debug.APIServer,
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
			time.Now(),
			logger,
			tracer,
			agentCacheUseCase,
			exportUseCases,
			cfg.AgentServer.Addr,
			cfg.Application.Debug.APIAgent,
			cfg.AgentServer.Token,
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
		infoCollector := metrics.NewSystemInfoCollector(
			logger,
			webAPIUseCases,
			storage,
			cfg.Application.Metric,
		)
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

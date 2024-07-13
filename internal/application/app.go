package application

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"hgnext/internal/adapters/agent"
	"hgnext/internal/adapters/files"
	"hgnext/internal/adapters/postgresql"
	"hgnext/internal/adapters/tmpdata"
	"hgnext/internal/controllers/apiserver"
	"hgnext/internal/controllers/workermanager"
	"hgnext/internal/metrics"
	agentUC "hgnext/internal/usecases/agent"
	"hgnext/internal/usecases/cleanup"
	"hgnext/internal/usecases/deduplicator"
	"hgnext/internal/usecases/export"
	"hgnext/internal/usecases/filelogic"
	"hgnext/internal/usecases/parsing"
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

	tmpStorage := tmpdata.New()

	storage, err := postgresql.New(ctx, cfg.PostgreSQLConnection, logger)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init postgres",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	fileStorage, err := files.New(cfg.FilePath, logger)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init file storage",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	agents, err := storage.Agents(ctx, false, false)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail load agents from storage",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	agentSystem, err := agent.New(agents)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init agent system",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	parsingUseCases := parsing.New(logger, storage, agentSystem, fileStorage)
	fileUseCases := filelogic.New(logger, storage, fileStorage)
	exportUseCases := export.New(logger, storage, fileStorage, agentSystem, tmpStorage)

	workersController := workermanager.New(
		logger,
		workermanager.NewBookParser(parsingUseCases, logger),
		workermanager.NewPageDownloader(parsingUseCases, logger),
		workermanager.NewHasher(fileUseCases, logger),
		workermanager.NewExporter(exportUseCases, logger),
	)

	webAPIUseCases := webapi.New(logger, workersController, storage, fileStorage)
	agentUseCases := agentUC.New(logger, agentSystem, storage)
	dededuplicateUseCases := deduplicator.New(logger, storage)
	cleanupUseCases := cleanup.New(logger, storage, fileStorage)

	apiController, err := apiserver.New(
		logger,
		cfg.WebServerAddr,
		cfg.ExternalWebServerAddr,
		parsingUseCases,
		webAPIUseCases,
		agentUseCases,
		exportUseCases,
		dededuplicateUseCases,
		cleanupUseCases,
		cfg.Debug,
		cfg.WebStaticDir,
		cfg.APIToken,
	)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail to create api server",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	err = metrics.RegisterSystemInfoCollector(logger, webAPIUseCases)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail to create system metric",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	asyncController := New(logger)
	asyncController.RegisterRunner(workersController)
	asyncController.RegisterRunner(apiController)

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

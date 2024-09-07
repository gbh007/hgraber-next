package server

import (
	"context"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"

	"hgnext/internal/adapters/agent"
	"hgnext/internal/adapters/agentFS"
	"hgnext/internal/adapters/files"
	"hgnext/internal/adapters/postgresql"
	"hgnext/internal/adapters/tmpdata"
	"hgnext/internal/controllers/apiagent"
	"hgnext/internal/controllers/apiserver"
	"hgnext/internal/controllers/workermanager"
	"hgnext/internal/entities"
	"hgnext/internal/metrics"
	agentUC "hgnext/internal/usecases/agent"
	"hgnext/internal/usecases/agentcache"
	"hgnext/internal/usecases/bookrequester"
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

	storage, err := postgresql.New(ctx, cfg.Storage.Connection, cfg.Storage.MaxConnections, logger)
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

	var fileStorage interface {
		Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
		Delete(ctx context.Context, fileID uuid.UUID) error
		Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
		IDs(ctx context.Context) ([]uuid.UUID, error)
	}

	switch {
	case cfg.Storage.FSAgentID != uuid.Nil:
		fileStorage = agentFS.New(cfg.Storage.FSAgentID, logger, agentSystem)

		logger.DebugContext(
			ctx, "use agent file storage",
			slog.String("agent_id", cfg.Storage.FSAgentID.String()),
		)

	case cfg.Storage.FilePath != "":
		fileStorage, err = files.New(cfg.Storage.FilePath, logger)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init local file storage",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		logger.DebugContext(
			ctx, "use local file storage",
			slog.String("path", cfg.Storage.FilePath),
		)

	default:
		logger.ErrorContext(
			ctx, "no configuration for file storage",
		)

		os.Exit(1)
	}

	bookRequestUseCases := bookrequester.New(logger, storage)
	parsingUseCases := parsing.New(logger, storage, agentSystem, fileStorage, bookRequestUseCases, cfg.Parsing.ParseBookTimeout)
	fileUseCases := filelogic.New(logger, storage, fileStorage)
	exportUseCases := export.New(logger, storage, fileStorage, agentSystem, tmpStorage, bookRequestUseCases)

	metricProvider := metrics.MetricProvider{}

	workersController := workermanager.New(
		logger,
		workermanager.NewBookParser(parsingUseCases, logger, tracer, cfg.Workers.Book, metricProvider),
		workermanager.NewPageDownloader(parsingUseCases, logger, tracer, cfg.Workers.Page, metricProvider),
		workermanager.NewHasher(fileUseCases, logger, tracer, cfg.Workers.Hasher, metricProvider),
		workermanager.NewExporter(exportUseCases, logger, tracer, cfg.Workers.Exporter, metricProvider),
	)
	asyncController.RegisterRunner(workersController)

	webAPIUseCases := webapi.New(logger, workersController, storage, fileStorage, bookRequestUseCases)
	agentUseCases := agentUC.New(logger, agentSystem, storage)
	dededuplicateUseCases := deduplicator.New(logger, storage, tracer)
	cleanupUseCases := cleanup.New(logger, tracer, storage, fileStorage)

	apiController, err := apiserver.New(
		logger,
		tracer,
		cfg.API,
		parsingUseCases,
		webAPIUseCases,
		agentUseCases,
		exportUseCases,
		dededuplicateUseCases,
		cleanupUseCases,
		cfg.Application.Debug,
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
			agentCacheUseCase,
			exportUseCases,
			cfg.AgentServer.Addr,
			cfg.Application.Debug,
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

	if cfg.Application.MetricTimeout > 0 {
		err = metrics.RegisterSystemInfoCollector(logger, webAPIUseCases, cfg.Application.MetricTimeout)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail to create system metric",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
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

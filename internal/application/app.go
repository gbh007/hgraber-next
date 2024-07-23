package application

import (
	"context"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"

	"hgnext/internal/adapters/agent"
	"hgnext/internal/adapters/agentFS"
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

	// TODO: использовать более подходящую проверку
	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" {
		err := initTrace(ctx)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init otel",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	tracer := otel.GetTracerProvider().Tracer("hgraber-next")

	tmpStorage := tmpdata.New()

	storage, err := postgresql.New(ctx, cfg.PostgreSQLConnection, logger)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init postgres",
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

	var fileStorage interface {
		Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
		Delete(ctx context.Context, fileID uuid.UUID) error
		Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
		IDs(ctx context.Context) ([]uuid.UUID, error)
	}

	switch {
	case cfg.FSAgentID != uuid.Nil:
		fileStorage = agentFS.New(cfg.FSAgentID, logger, agentSystem)

		logger.DebugContext(
			ctx, "use agent file storage",
			slog.String("agent_id", cfg.FSAgentID.String()),
		)

	case cfg.FilePath != "":
		fileStorage, err = files.New(cfg.FilePath, logger)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init local file storage",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		logger.DebugContext(
			ctx, "use local file storage",
			slog.String("path", cfg.FilePath),
		)

	default:
		logger.ErrorContext(
			ctx, "no configuration for file storage",
		)

		os.Exit(1)
	}

	parsingUseCases := parsing.New(logger, storage, agentSystem, fileStorage, cfg.Handle.ParseBookTimeout)
	fileUseCases := filelogic.New(logger, storage, fileStorage)
	exportUseCases := export.New(logger, storage, fileStorage, agentSystem, tmpStorage)

	workersController := workermanager.New(
		logger,
		workermanager.NewBookParser(parsingUseCases, logger, tracer),
		workermanager.NewPageDownloader(parsingUseCases, logger, tracer),
		workermanager.NewHasher(fileUseCases, logger, tracer),
		workermanager.NewExporter(exportUseCases, logger, tracer),
	)

	webAPIUseCases := webapi.New(logger, workersController, storage, fileStorage)
	agentUseCases := agentUC.New(logger, agentSystem, storage)
	dededuplicateUseCases := deduplicator.New(logger, storage, tracer)
	cleanupUseCases := cleanup.New(logger, tracer, storage, fileStorage)

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

	if cfg.MetricTimeout > 0 {
		err = metrics.RegisterSystemInfoCollector(logger, webAPIUseCases, cfg.MetricTimeout)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail to create system metric",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
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

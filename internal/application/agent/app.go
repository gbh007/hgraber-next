package agent

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gbh007/hgraber-next/internal/agent/config"
	"github.com/gbh007/hgraber-next/internal/controllers/api"
	"github.com/gbh007/hgraber-next/internal/controllers/async"
	"github.com/gbh007/hgraber-next/internal/controllers/debugserver"
	"github.com/gbh007/hgraber-next/internal/dataproviders/datafs"
	"github.com/gbh007/hgraber-next/internal/dataproviders/importfs"
	"github.com/gbh007/hgraber-next/internal/dataproviders/loader"
	"github.com/gbh007/hgraber-next/internal/dataproviders/masterapi"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric"
	"github.com/gbh007/hgraber-next/internal/dataproviders/storage"
	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/domain/hgraber"
	"github.com/gbh007/hgraber-next/internal/pkg"
	agentUC "github.com/gbh007/hgraber-next/internal/usecases/agent"
	"github.com/gbh007/hgraber-next/internal/usecases/highway"
	"github.com/gbh007/hgraber-next/internal/usecases/importapi"
	"github.com/gbh007/hgraber-next/internal/usecases/importdeduplicator"
	"go.opentelemetry.io/otel"
)

type Async interface {
	RegisterRunner(runner async.Runner)
	RegisterAfterStop(handler func())
}

type ParserInit[T any] func(
	ctx context.Context,
	logger *slog.Logger,
	cfg config.AgentConfig[T],
	async Async,
	metricProvider *metric.MetricProvider,
) ([]hgraber.Parser, error)

func Serve[T any](ctx context.Context, parserInit ParserInit[T], defaultParsers func() *T) {
	cfg, needScan, err := parseConfig[T](defaultParsers)
	if err != nil {
		// Поскольку на этот момент нет ни логгера ни вообще ничего то выкидываем панику.
		panic(err)
	}

	logger := initLogger(cfg)
	logger.InfoContext(ctx, "initializing system")

	metricProvider, err := metric.New(metric.Config{
		ServiceName:    cfg.Application.ServiceName,
		Type:           metric.AgentSystemType,
		WithGo:         true,
		WithVersion:    true,
		WithFS:         true,
		WithServer:     false,
		WithDB:         true,
		WithHTTPServer: true,
		WithAgent:      true,
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
		err := initTrace(
			ctx,
			cfg.Application.TraceEndpoint,
			cfg.Application.ServiceName,
		)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init otel",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	tracer := otel.GetTracerProvider().Tracer("hgraber-next-agent")

	async := async.New(logger)

	var (
		importStorage   api.ImportUseCases
		fileStorage     api.FileUseCases
		agentUseCases   api.ParsingUseCases
		highwayUseCases api.HighwayUseCases

		importStorageRaw *importfs.Storage
		dbRaw            *storage.Storage
		mAPI             *masterapi.Client
	)

	parsers, err := parserInit(ctx, logger, cfg, async, metricProvider)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init parsers",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	if len(parsers) > 0 {
		loader := loader.New(parsers, metricProvider)
		agentUseCases = agentUC.New(logger, loader)

		logger.DebugContext(
			ctx, "use parsing",
		)
	}

	if cfg.FSBase.ImportPath != "" {
		importStorageRaw, err = importfs.New(
			logger,
			metricProvider,
			cfg.FSBase.ImportPath,
			cfg.FSBase.ImportLimitOnFolder,
			cfg.Application.UseUnsafeCloser,
		)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init import fs",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		importStorage = importStorageRaw

		logger.DebugContext(
			ctx, "use local import storage",
			slog.String("path", cfg.FSBase.ImportPath),
		)
	}

	if cfg.FSBase.FilePath != "" {
		fileStorage, err = datafs.New(logger, metricProvider, cfg.FSBase.FilePath)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init data fs",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		logger.DebugContext(
			ctx, "use local file storage",
			slog.String("path", cfg.FSBase.FilePath),
		)
	}

	if cfg.Sqlite.FilePath != "" {
		dbRaw, err = storage.New(ctx, logger, metricProvider, cfg.Sqlite.FilePath)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init db",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	if cfg.ZipScanner.MasterAddr != "" {
		mAPI, err = masterapi.New(cfg.ZipScanner.MasterAddr, cfg.ZipScanner.MasterToken)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init master api",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	if needScan {
		if dbRaw == nil || importStorageRaw == nil || mAPI == nil {
			logger.ErrorContext(ctx, "invalid scan dependencies")

			os.Exit(1)
		}

		err = importdeduplicator.New(logger, importStorageRaw, dbRaw, mAPI).ScanZips(ctx)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail scan zips",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		return
	}

	if cfg.FSBase.EnableDeduplication && dbRaw != nil && importStorageRaw != nil {
		importStorage = importapi.New(logger, dbRaw, importStorageRaw)

		logger.DebugContext(ctx, "use export deduplication")
	}

	if fileStorage != nil && cfg.Highway.PrivateKey != "" && cfg.Highway.TokenLifetime > 0 {
		tokenizer, err := entities.NewSimpleHighwayTokenizer(cfg.Highway.PrivateKey)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init highway tokenizer",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		highwayUseCases = highway.New(
			tokenizer,
			cfg.Highway.TokenLifetime,
			fileStorage,
		)
	}

	parserNames := pkg.Map(parsers, func(parser hgraber.Parser) string {
		return parser.Name()
	})

	if cfg.API.Addr != "" {
		apiController, err := api.New(
			cfg.API,
			time.Now(),
			logger,
			tracer,
			agentUseCases,
			importStorage,
			fileStorage,
			highwayUseCases,
			parserNames,
			metricProvider,
		)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init api controller",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		async.RegisterRunner(apiController)
	}

	if cfg.DebugServer.Addr != "" && agentUseCases != nil {
		async.RegisterRunner(debugserver.New[T](
			cfg,
			logger,
			agentUseCases,
		))
	}

	logger.InfoContext(ctx, "application start")
	defer logger.InfoContext(ctx, "application stop")

	err = async.Serve(ctx)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail serve",
			slog.Any("error", err),
		)

		os.Exit(1)
	}
}

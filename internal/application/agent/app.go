package agent

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
	"github.com/gbh007/hgraber-next/internal/controllers/agentapi"
	"github.com/gbh007/hgraber-next/internal/controllers/agentdebugserver"
	"github.com/gbh007/hgraber-next/internal/controllers/async"
	"github.com/gbh007/hgraber-next/internal/dataproviders/agentdb"
	"github.com/gbh007/hgraber-next/internal/dataproviders/agentimportfs"
	"github.com/gbh007/hgraber-next/internal/dataproviders/localfs"
	"github.com/gbh007/hgraber-next/internal/dataproviders/masterapi"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parserloader"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/internal/usecases/agenthighway"
	"github.com/gbh007/hgraber-next/internal/usecases/agentimportapiuc"
	"github.com/gbh007/hgraber-next/internal/usecases/agentimportdeduplicatoruc"
	"github.com/gbh007/hgraber-next/internal/usecases/agentparseruc"
	"go.opentelemetry.io/otel"
)

type Async interface {
	RegisterRunner(runner async.Runner)
	RegisterAfterStop(handler func())
}

type ParserInit = func(
	ctx context.Context,
	logger *slog.Logger,
	cfg agnconfig.AgentConfig,
	async Async,
	metricProvider *metric.MetricProvider,
) ([]parserdto.Parser, error)

func Serve(ctx context.Context, parserInit ParserInit) {
	cfg, needScan, err := parseConfig()
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

	tracer := otel.GetTracerProvider().Tracer("hgraber-next-agent")

	async := async.New(logger)

	var (
		importStorage   agentapi.ImportUseCases
		fileStorage     agentapi.FileUseCases
		agentUseCases   agentapi.ParsingUseCases
		highwayUseCases agentapi.HighwayUseCases

		importStorageRaw *agentimportfs.Storage
		dbRaw            *agentdb.Storage
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
		loader := parserloader.New(parsers, metricProvider)
		agentUseCases = agentparseruc.New(logger, loader)

		logger.DebugContext(
			ctx, "use parsing",
		)
	}

	if cfg.FSBase.ImportPath != "" {
		importStorageRaw, err = agentimportfs.New(
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
		fileStorage, err = localfs.New(cfg.FSBase.FilePath, logger)
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
		dbRaw, err = agentdb.New(ctx, logger, metricProvider, cfg.Sqlite.FilePath)
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

		err = agentimportdeduplicatoruc.New(logger, importStorageRaw, dbRaw, mAPI).ScanZips(ctx)
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
		importStorage = agentimportapiuc.New(logger, dbRaw, importStorageRaw)

		logger.DebugContext(ctx, "use export deduplication")
	}

	if fileStorage != nil && cfg.Highway.PrivateKey != "" && cfg.Highway.TokenLifetime > 0 {
		tokenizer, err := serveragentdto.NewSimpleHighwayTokenizer(cfg.Highway.PrivateKey)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init highway tokenizer",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		highwayUseCases = agenthighway.New(
			tokenizer,
			cfg.Highway.TokenLifetime,
			fileStorage,
		)
	}

	parserNames := pkg.Map(parsers, func(parser parserdto.Parser) string {
		return parser.Name()
	})

	if cfg.API.Addr != "" {
		apiController, err := agentapi.New(
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
		async.RegisterRunner(agentdebugserver.New(
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

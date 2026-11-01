package agent

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/controllers/agentapi"
	"github.com/gbh007/hgraber-next/internal/controllers/agentdebugserver"
	"github.com/gbh007/hgraber-next/internal/controllers/async"
	"github.com/gbh007/hgraber-next/internal/dataproviders/agentdb"
	"github.com/gbh007/hgraber-next/internal/dataproviders/agentimportfs"
	"github.com/gbh007/hgraber-next/internal/dataproviders/agentparsecache"
	"github.com/gbh007/hgraber-next/internal/dataproviders/localfs"
	"github.com/gbh007/hgraber-next/internal/dataproviders/masterapi"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parserloader"
	"github.com/gbh007/hgraber-next/internal/dataproviders/request"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/internal/usecases/agenthighway"
	"github.com/gbh007/hgraber-next/internal/usecases/agentimportapiuc"
	"github.com/gbh007/hgraber-next/internal/usecases/agentimportdeduplicatoruc"
	"github.com/gbh007/hgraber-next/internal/usecases/agentparseruc"
)

func InitAndServe(parserInits map[string]parserdto.NewParser) {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	err := Serve(ctx, parserInits)
	if err != nil {
		slog.Default().Error(err.Error())

		//nolint:revive,gocritic // так и задумано
		os.Exit(1)
	}
}

func Serve(ctx context.Context, parserInits map[string]parserdto.NewParser) error {
	cfg, needScan, err := parseConfig()
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	logger := initLogger(cfg)
	slog.SetDefault(logger)

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
		return fmt.Errorf("fail init metrics: %w", err)
	}

	asyncCnt := async.New(logger)

	var (
		importStorage   agentapi.ImportUseCases
		fileStorage     agentapi.FileUseCases
		agentUseCases   agentapi.ParsingUseCases
		highwayUseCases agentapi.HighwayUseCases

		importStorageRaw *agentimportfs.Storage
		dbRaw            *agentdb.Storage
		mAPI             *masterapi.Client

		cache request.Cache
	)

	if cfg.Parsers.Cache.Enabled {
		wc, err := agentparsecache.New(
			cfg.Parsers.Cache.Path,
			logger,
			cfg.Parsers.Cache.TTL,
			cfg.Parsers.Cache.CleanInterval,
			metricProvider,
		)
		if err != nil {
			return fmt.Errorf("create web cache: %w", err)
		}

		asyncCnt.RegisterRunner(wc)
		cache = wc
	}

	parsers := initParsers(
		logger,
		cfg.Application.ClientTimeout,
		cfg.Parsers,
		cache,
		metricProvider,
		cfg.Log.RequestErrors,
		cfg.Parsers.UserAgent,
		parserInits,
	)

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
			return fmt.Errorf("fail init import fs: %w", err)
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
			return fmt.Errorf("fail init data fs: %w", err)
		}

		logger.DebugContext(
			ctx, "use local file storage",
			slog.String("path", cfg.FSBase.FilePath),
		)
	}

	if cfg.Sqlite.FilePath != "" {
		dbRaw, err = agentdb.New(ctx, logger, metricProvider, cfg.Sqlite.FilePath)
		if err != nil {
			return fmt.Errorf("fail init db: %w", err)
		}
	}

	if cfg.ZipScanner.MasterAddr != "" {
		mAPI, err = masterapi.New(cfg.ZipScanner.MasterAddr, cfg.ZipScanner.MasterToken)
		if err != nil {
			return fmt.Errorf("fail init master api client: %w", err)
		}
	}

	if needScan {
		if dbRaw == nil || importStorageRaw == nil || mAPI == nil {
			return errors.New("invalid scan dependencies")
		}

		err = agentimportdeduplicatoruc.New(logger, importStorageRaw, dbRaw, mAPI).ScanZips(ctx)
		if err != nil {
			return fmt.Errorf("fail scan zips: %w", err)
		}

		return nil
	}

	if cfg.FSBase.EnableDeduplication && dbRaw != nil && importStorageRaw != nil {
		importStorage = agentimportapiuc.New(logger, dbRaw, importStorageRaw)

		logger.DebugContext(ctx, "use export deduplication")
	}

	if fileStorage != nil && cfg.Highway.PrivateKey != "" && cfg.Highway.TokenLifetime > 0 {
		tokenizer, err := serveragentdto.NewSimpleHighwayTokenizer(cfg.Highway.PrivateKey)
		if err != nil {
			return fmt.Errorf("fail init highway tokenizer: %w", err)
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
			agentUseCases,
			importStorage,
			fileStorage,
			highwayUseCases,
			parserNames,
			metricProvider,
		)
		if err != nil {
			return fmt.Errorf("fail init api controller: %w", err)
		}

		asyncCnt.RegisterRunner(apiController)
	}

	if cfg.DebugServer.Addr != "" && agentUseCases != nil {
		asyncCnt.RegisterRunner(agentdebugserver.New(
			cfg,
			logger,
			agentUseCases,
		))
	}

	logger.InfoContext(ctx, "application start")
	defer logger.InfoContext(ctx, "application stop")

	err = asyncCnt.Serve(ctx)
	if err != nil {
		return fmt.Errorf("fail serve: %w", err)
	}

	return nil
}

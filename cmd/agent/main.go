package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/gbh007/hgraber-next/internal/application/agent"
	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
	"github.com/gbh007/hgraber-next/internal/dataproviders/agentparsecache"
	"github.com/gbh007/hgraber-next/internal/dataproviders/metric"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parserloader"
	"github.com/gbh007/hgraber-next/internal/dataproviders/request"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	agent.Serve(ctx, func(
		ctx context.Context,
		logger *slog.Logger,
		cfg agnconfig.AgentConfig[agnconfig.AgentParsers],
		async agent.Async,
		metricProvider *metric.MetricProvider,
	) ([]parsing.Parser, error) {
		if cfg.Parsers == nil {
			logger.DebugContext(ctx, "nil parser config, skipping")

			return []parsing.Parser{}, nil
		}

		var cache request.Cache

		if cfg.Parsers.Cache.Enabled {
			wc, err := agentparsecache.New(
				cfg.Parsers.Cache.Path,
				logger,
				cfg.Parsers.Cache.TTL,
				cfg.Parsers.Cache.CleanInterval,
				metricProvider,
			)
			if err != nil {
				return nil, fmt.Errorf("create web cache: %w", err)
			}

			async.RegisterRunner(wc)
			cache = wc
		}

		return parserloader.NewDefaultParsers(
			logger,
			cfg.Application.ClientTimeout,
			*cfg.Parsers,
			cache,
			metricProvider,
			cfg.Log.RequestErrors,
			cfg.Parsers.UserAgent,
		), nil
	}, agnconfig.AgentDefaultParsers)
}

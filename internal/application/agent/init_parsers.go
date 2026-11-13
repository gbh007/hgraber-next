package agent

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
	"github.com/gbh007/hgraber-next/internal/dataproviders/request"
	"github.com/samber/lo"
	"golang.org/x/net/proxy"
)

func initParsers(
	logger *slog.Logger,
	timeout time.Duration,
	cfg agnconfig.AgentParsers,
	cache request.Cache,
	metricProvider parserdto.MetricProvider,
	logErrors bool,
	defaultUserAgent string,
	parserInits map[string]parserdto.NewParser,
) []parserdto.Parser {
	var (
		withProxy      map[string]struct{}
		proxyRequester *request.Requester
	)

	if cfg.Proxy.Host != "" {
		withProxy = lo.SliceToMap(cfg.Proxy.Enabled, func(s string) (string, struct{}) { return s, struct{}{} })

		var auth *proxy.Auth

		if cfg.Proxy.User != "" {
			auth = &proxy.Auth{
				User:     cfg.Proxy.User,
				Password: cfg.Proxy.Pass,
			}
		}

		dialer, err := proxy.SOCKS5("tcp", cfg.Proxy.Host, auth, proxy.Direct)
		if err != nil {
			panic(err)
		}

		proxyRequester = request.New(
			logger,
			timeout,
			&http.Transport{
				//nolint:staticcheck // нет другого варианта
				Dial: dialer.Dial,
			},
			logErrors,
			cache,
			defaultUserAgent,
		)
	}

	requester := request.New(logger, timeout, nil, logErrors, cache, defaultUserAgent)

	parsers := make([]parserdto.Parser, 0, len(cfg.Enabled))

	for _, code := range cfg.Enabled {
		rq := requester

		if _, ok := withProxy[code]; ok {
			rq = proxyRequester
		}

		parserInit, ok := parserInits[code]
		if !ok {
			logger.Warn(
				"unknown parser code",
				slog.String("code", code),
			)

			continue
		}

		parsers = append(parsers, parserInit(rq, metricProvider))
	}

	return parsers
}

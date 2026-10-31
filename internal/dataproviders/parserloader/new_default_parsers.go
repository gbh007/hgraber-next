package parserloader

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/common"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/hgraber_local"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/mock"
	"github.com/gbh007/hgraber-next/internal/dataproviders/request"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"golang.org/x/net/proxy"
)

func NewDefaultParsers(
	logger *slog.Logger,
	timeout time.Duration,
	cfg agnconfig.AgentParsers,
	cache request.Cache,
	metricProvider metricProvider,
	logErrors bool,
	defaultUserAgent string,
) []parsing.Parser {
	var (
		withProxy      map[string]struct{}
		proxyRequester *request.Requester
		parserInits    map[string]common.NewParser
	)

	// FIXME: поработать над этим блоком пока прибиваем гвоздями
	parserInits = map[string]common.NewParser{
		"mock":          mock.New,
		"hgraber_local": hgraber_local.New,
	}

	if cfg.Proxy.Host != "" {
		withProxy = pkg.SliceToMap(cfg.Proxy.Enabled, func(s string) (string, struct{}) { return s, struct{}{} })

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
				Dial: dialer.Dial,
			},
			logErrors,
			cache,
			defaultUserAgent,
		)
	}

	requester := request.New(logger, timeout, nil, logErrors, cache, defaultUserAgent)

	parsers := make([]parsing.Parser, 0, len(cfg.Enabled))

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

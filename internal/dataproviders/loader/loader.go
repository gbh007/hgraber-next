package loader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/hgraber"
	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/hgraber_local"
	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/mock"
	"github.com/gbh007/hgraber-next/internal/dataproviders/request"
)

type metricProvider interface {
	RegisterParserActionTime(action, parserName string, d time.Duration)
}

type Loader struct {
	parsers        []hgraber.Parser
	metricProvider metricProvider
}

func NewDefaultParsers(
	logger *slog.Logger,
	hgToken string,
	timeout time.Duration,
	enabledParsers []string,
	cache request.Cache,
	logErrors bool,
	defaultUserAgent string,
) []hgraber.Parser {
	requester := request.New(logger, timeout, nil, logErrors, cache, defaultUserAgent)

	parsers := make([]hgraber.Parser, 0, len(enabledParsers))

	for _, code := range enabledParsers {
		var p hgraber.Parser

		switch code {
		case "mock":
			p = mock.New(requester)

		case "hgraber_local":
			p = hgraber_local.New(requester, hgToken)

		default:
			logger.Warn(
				"unknown parser code",
				slog.String("code", code),
			)

			continue
		}

		parsers = append(parsers, p)
	}

	return parsers
}

func New(
	parsers []hgraber.Parser,
	metricProvider metricProvider,
) *Loader {
	return &Loader{
		parsers:        parsers,
		metricProvider: metricProvider,
	}
}

func (l *Loader) getParser(u url.URL) (hgraber.Parser, error) {
	for _, p := range l.parsers {
		if p.CanParse(u) {
			return p, nil
		}
	}

	return nil, hgraber.InvalidLinkError
}

func (l *Loader) HasParser(ctx context.Context, u url.URL) (bool, error) {
	_, err := l.getParser(u)
	if errors.Is(err, hgraber.InvalidLinkError) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("get parser: %w", err)
	}

	return true, nil
}

func (l *Loader) Load(ctx context.Context, u url.URL) (hgraber.BookParser, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return nil, fmt.Errorf("get parser: %w", err)
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_load", p.Name(), time.Since(startAt))
	}()

	bookParser, err := p.Load(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}

	return bookParser, nil
}

func (l *Loader) LoadImage(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error) {
	startAt := time.Now()

	p, err := l.getParser(bookUrl)
	if err != nil {
		return nil, fmt.Errorf("get parser: %w", err)
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_load_image", p.Name(), time.Since(startAt))
	}()

	data, err := p.LoadImage(ctx, u, bookUrl)
	if err != nil {
		return nil, fmt.Errorf("load image: %w", err)
	}

	return data, nil
}

func (l *Loader) AllBooks(ctx context.Context, u url.URL) ([]string, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return nil, err
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_all_books", p.Name(), time.Since(startAt))
	}()

	data, err := p.AllBooks(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("load books: %w", err)
	}

	return data, nil
}

func (l *Loader) HProxyList(ctx context.Context, u url.URL) (entities.HProxyList, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return entities.HProxyList{}, err
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_hproxy_list", p.Name(), time.Since(startAt))
	}()

	data, err := p.HProxyList(ctx, u)
	if err != nil {
		return entities.HProxyList{}, fmt.Errorf("load hproxy list: %w", err)
	}

	return data, nil
}

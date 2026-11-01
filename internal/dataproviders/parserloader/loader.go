package parserloader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/gbh007/hgraber-next/external/parserdto"
)

type metricProvider interface {
	RegisterParserActionTime(action, parserName string, d time.Duration)
}

type Loader struct {
	parsers        []parserdto.Parser
	metricProvider metricProvider
}

func New(
	parsers []parserdto.Parser,
	metricProvider metricProvider,
) *Loader {
	return &Loader{
		parsers:        parsers,
		metricProvider: metricProvider,
	}
}

func (l *Loader) getParser(u url.URL) (parserdto.Parser, error) {
	for _, p := range l.parsers {
		if p.CanParse(u) {
			return p, nil
		}
	}

	return nil, parserdto.InvalidLinkError
}

func (l *Loader) HasParser(ctx context.Context, u url.URL) (bool, error) {
	_, err := l.getParser(u)
	if errors.Is(err, parserdto.InvalidLinkError) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("get parser: %w", err)
	}

	return true, nil
}

func (l *Loader) Book(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("get parser: %w", err)
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_load", p.Name(), time.Since(startAt))
	}()

	bookParser, err := p.Book(ctx, u, pageLimit)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("load: %w", err)
	}

	return bookParser, nil
}

func (l *Loader) Image(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error) {
	startAt := time.Now()

	p, err := l.getParser(bookUrl)
	if err != nil {
		return nil, fmt.Errorf("get parser: %w", err)
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_load_image", p.Name(), time.Since(startAt))
	}()

	data, err := p.Image(ctx, u, bookUrl)
	if err != nil {
		return nil, fmt.Errorf("load image: %w", err)
	}

	return data, nil
}

func (l *Loader) List(ctx context.Context, u url.URL) (parserdto.BookList, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return parserdto.BookList{}, err
	}

	defer func() {
		l.metricProvider.RegisterParserActionTime("loader_hproxy_list", p.Name(), time.Since(startAt))
	}()

	data, err := p.List(ctx, u)
	if err != nil {
		return parserdto.BookList{}, fmt.Errorf("load hproxy list: %w", err)
	}

	return data, nil
}

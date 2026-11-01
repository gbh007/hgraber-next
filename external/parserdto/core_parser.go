package parserdto

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Проверка соответствия базового типа
var (
	_ Parser = (*CoreParser)(nil)
)

type Requester interface {
	Request(ctx context.Context, URL string, headers http.Header) (io.ReadCloser, error)
	RequestString(ctx context.Context, URL string) (string, error)
	RequestStringWithRedirect(ctx context.Context, URL string) (string, string, error)
	RequestPost(ctx context.Context, u string, headers http.Header, body io.Reader) ([]byte, error)
}

type MetricProvider interface {
	RegisterParserActionTime(action, parserName string, d time.Duration)
}

type NewParser = func(Requester, MetricProvider) Parser

type CoreParser struct {
	Requester Requester
	prefixes  []string
	name      string
}

func NewCoreParser(requester Requester, prefixes []string, name string) CoreParser {
	return CoreParser{
		Requester: requester,
		prefixes:  prefixes,
		name:      name,
	}
}

func (cp CoreParser) Name() string {
	return cp.name
}

func (cp CoreParser) Book(ctx context.Context, u url.URL, pageLimit *int) (Book, error) {
	return Book{}, errors.New("unimplemented in core parser")
}

func (cp CoreParser) CanParse(u url.URL) bool {
	for _, prefix := range cp.prefixes {
		if strings.HasPrefix(u.String(), prefix) {
			return true
		}
	}

	return false
}

func (cp CoreParser) Image(ctx context.Context, u, bookUrl url.URL) (io.ReadCloser, error) {
	data, err := cp.Requester.Request(ctx, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cp CoreParser) List(ctx context.Context, u url.URL) (BookList, error) {
	return BookList{}, nil
}

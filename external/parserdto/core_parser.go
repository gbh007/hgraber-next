package parserdto

import (
	"context"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
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

func (cp CoreParser) Load(ctx context.Context, u url.URL) (BookParser, error) {
	return nil, fmt.Errorf("unimplemented in core parser")
}

func (cp CoreParser) CanParse(u url.URL) bool {
	for _, prefix := range cp.prefixes {
		if strings.HasPrefix(u.String(), prefix) {
			return true
		}
	}

	return false
}

func (cp CoreParser) AllBooks(ctx context.Context, u url.URL) ([]string, error) {
	return nil, nil
}

func (cp CoreParser) LoadImage(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error) {
	data, err := cp.Requester.Request(ctx, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cp CoreParser) HProxyList(ctx context.Context, u url.URL) (BookList, error) {
	return BookList{}, nil
}

func TrimLastSlash(URL string, count int) string {
	c := 0

	ind := strings.LastIndexFunc(URL, func(r rune) bool {
		if r != rune('/') {
			return false
		}
		c++
		return c == count
	})

	return URL[:ind]
}

func OneMatch(rgx *regexp.Regexp, raw string) (string, bool) {
	res := rgx.FindAllStringSubmatch(raw, -1)
	if len(res) < 1 || len(res[0]) != 2 {
		return "", false
	}

	return strings.TrimSpace(html.UnescapeString(res[0][1])), true
}

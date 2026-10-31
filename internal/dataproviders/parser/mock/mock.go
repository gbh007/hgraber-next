package mock

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/common"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
)

// Проверка соответствия базового типа
var (
	_ parsing.Parser = (*Parser)(nil)

	ParserError = errors.New("parser mock")
)

type Parser struct {
	common.CoreParser
}

func New(r common.Requester, _ common.MetricProvider) parsing.Parser {
	return &Parser{
		CoreParser: common.NewCoreParser(r, []string{
			"http://localhost",
		}, "mock"),
	}
}

func (p *Parser) Load(ctx context.Context, u url.URL) (parsing.BookParser, error) {
	bookParser := BookParser{
		url: u.String(),
	}

	if len(bookParser.url) > 1 && bookParser.url[len(bookParser.url)-1] == '/' {
		bookParser.url = bookParser.url[:len(bookParser.url)-1]
	}

	body, err := p.Requester.RequestString(ctx, u.String())
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ParserError, err)
	}

	bookParser.body = body

	return common.NewLegacyParserAdapter(bookParser, u), nil
}

type BookParser struct {
	url, body string
}

func (p BookParser) Pages(ctx context.Context) ([]parsing.RawPage, error) {
	result := make([]parsing.RawPage, 0)

	rp := `(?sm)` + regexp.QuoteMeta(`<a href="`) + `(.+?)\.(.+?)` + regexp.QuoteMeta(`">`)
	for i, name := range regexp.MustCompile(rp).FindAllStringSubmatch(p.body, -1) {
		if len(name) > 1 {
			result = append(result, parsing.RawPage{
				PageNumber: i + 1,
				URL:        p.url + "/" + strings.TrimSpace(name[1]) + "." + strings.TrimSpace(name[2]),
				Ext:        name[2],
			})
		}
	}

	return result, nil
}

func (p BookParser) Name(ctx context.Context) (string, error) {
	return "mock name", nil
}

func (p BookParser) Tags(ctx context.Context) ([]string, error) {
	return []string{"mock Tags"}, nil
}

func (p BookParser) Authors(ctx context.Context) ([]string, error) {
	return []string{"mock Authors"}, nil
}

func (p BookParser) Characters(ctx context.Context) ([]string, error) {
	return []string{"mock Characters"}, nil
}

func (p BookParser) Languages(ctx context.Context) ([]string, error) {
	return []string{"mock Languages"}, nil
}

func (p BookParser) Categories(ctx context.Context) ([]string, error) {
	return []string{"mock Categories"}, nil
}

func (p BookParser) Parodies(ctx context.Context) ([]string, error) {
	return []string{"mock Parodies"}, nil
}

func (p BookParser) Groups(ctx context.Context) ([]string, error) {
	return []string{"mock Groups"}, nil
}

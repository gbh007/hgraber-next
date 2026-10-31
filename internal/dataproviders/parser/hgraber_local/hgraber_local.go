package hgraber_local

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gbh007/hgraber-next/internal/dataproviders/parser/common"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
)

// Проверка соответствия базового типа
var (
	_ parsing.Parser = (*Parser)(nil)

	ParserError = errors.New("parser hgraber_v4")
)

type Parser struct {
	common.CoreParser
	token string
}

func New(r common.Requester, _ common.MetricProvider) parsing.Parser {
	return &Parser{
		CoreParser: common.NewCoreParser(r, []string{
			"hg4://",
			"hgraber4://",
		}, "hgraber_local"),
		token: os.Getenv("HGRABER_PARSER_HG4_TOKEN"),
	}
}

func (p *Parser) Load(ctx context.Context, u url.URL) (parsing.BookParser, error) {
	bookParser := BookParser{}

	id, err := strconv.Atoi(strings.TrimPrefix(u.Path, "/"))
	if err != nil {
		return nil, fmt.Errorf("%w: parse book id: %w", ParserError, err)
	}

	u.Scheme = "http"
	u.Path = "/api/book"

	requestBody, err := json.Marshal(map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: marshal request body: %w", ParserError, err)
	}

	raw, err := p.Requester.RequestPost(ctx, u.String(), p.headers(), bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("%w: request: %w", ParserError, err)
	}

	err = json.Unmarshal(raw, &bookParser.body)
	if err != nil {
		return nil, fmt.Errorf("%w: unmarshal: %w", ParserError, err)
	}

	return common.NewLegacyParserAdapter(bookParser, u), nil
}

func (p *Parser) headers() http.Header {
	headers := make(http.Header, 1)

	headers.Set("x-token", p.token)

	return headers
}

func (p *Parser) LoadImage(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error) {
	data, err := p.Requester.Request(ctx, u.String(), p.headers())
	if err != nil {
		return nil, err
	}

	return data, nil
}

type BookParser struct {
	body bookDetailInfo
}

func (p BookParser) Name(ctx context.Context) (string, error) {
	return p.body.Name, nil
}

func (p BookParser) Pages(ctx context.Context) ([]parsing.RawPage, error) {
	result := make([]parsing.RawPage, len(p.body.Pages))

	for i, p := range p.body.Pages {
		result[i] = parsing.RawPage{
			PageNumber: p.PageNumber,
			URL:        p.PreviewURL,
			Ext:        path.Ext(p.PreviewURL), // FIXME: а сработает?
		}
	}

	return result, nil
}

func (p BookParser) Tags(_ context.Context) ([]string, error) {
	return p.parseTags("Тэги"), nil
}

func (p BookParser) Authors(_ context.Context) ([]string, error) {
	return p.parseTags("Авторы"), nil
}

func (p BookParser) Languages(_ context.Context) ([]string, error) {
	return p.parseTags("Языки"), nil
}

func (p BookParser) Parodies(_ context.Context) ([]string, error) {
	return p.parseTags("Пародии"), nil
}

func (p BookParser) Characters(_ context.Context) ([]string, error) {
	return p.parseTags("Персонажи"), nil
}

func (p BookParser) Categories(_ context.Context) ([]string, error) {
	return p.parseTags("Категории"), nil
}

func (p BookParser) Groups(_ context.Context) ([]string, error) {
	return p.parseTags("Группы"), nil
}

func (p BookParser) parseTags(name string) []string {
	for _, attr := range p.body.Attributes {
		if attr.Name == name {
			return attr.Values
		}
	}

	return nil
}

type bookDetailInfo struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`

	PreviewURL string `json:"preview_url,omitempty"`

	ParsedName bool   `json:"parsed_name"`
	Name       string `json:"name"`

	ParsedPage        bool    `json:"parsed_page"`
	PageCount         int     `json:"page_count"`
	PageLoadedPercent float64 `json:"page_loaded_percent"`

	Rating int `json:"rating"`

	Attributes []bookDetailAttributeInfo `json:"attributes,omitempty"`
	Pages      []bookDetailPagePreview   `json:"pages,omitempty"`
}

type bookDetailAttributeInfo struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type bookDetailPagePreview struct {
	PageNumber int    `json:"page_number"`
	PreviewURL string `json:"preview_url,omitempty"`
	Rating     int    `json:"rating"`
}

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

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/external/parserdto/legacyadapter"
)

// Проверка соответствия базового типа
var (
	_ parserdto.Parser = (*Parser)(nil)

	errParser = errors.New("parser hgraber_v4")
)

type Parser struct {
	parserdto.CoreParser

	token string
}

func New(r parserdto.Requester, _ parserdto.MetricProvider) parserdto.Parser {
	return &Parser{
		CoreParser: parserdto.NewCoreParser(r, []string{
			"hg4://",
			"hgraber4://",
		}, "hgraber_local"),
		token: os.Getenv("HGRABER_PARSER_HG4_TOKEN"),
	}
}

func (p *Parser) Book(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error) {
	bookParser := BookParser{}

	id, err := strconv.Atoi(strings.TrimPrefix(u.Path, "/"))
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("%w: parse book id: %w", errParser, err)
	}

	u.Scheme = "http"
	u.Path = "/api/book"

	requestBody, err := json.Marshal(map[string]any{
		"id": id,
	})
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("%w: marshal request body: %w", errParser, err)
	}

	raw, err := p.Requester.RequestPost(ctx, u.String(), p.headers(), bytes.NewReader(requestBody))
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("%w: request: %w", errParser, err)
	}

	err = json.Unmarshal(raw, &bookParser.body)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("%w: unmarshal: %w", errParser, err)
	}

	return legacyadapter.Convert(ctx, u, bookParser)
}

func (p *Parser) headers() http.Header {
	headers := make(http.Header, 1)

	headers.Set("x-token", p.token)

	return headers
}

func (p *Parser) Image(ctx context.Context, u, bookUrl url.URL) (io.ReadCloser, error) {
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

func (p BookParser) Pages(ctx context.Context) ([]legacyadapter.Page, error) {
	result := make([]legacyadapter.Page, len(p.body.Pages))

	for i, p := range p.body.Pages {
		result[i] = legacyadapter.Page{
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

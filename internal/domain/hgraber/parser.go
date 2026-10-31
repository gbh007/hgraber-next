package hgraber

import (
	"context"
	"errors"
	"io"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

var (
	InvalidLinkError      = errors.New("invalid link")
	UnknownAttributeError = errors.New("unknown attribute")
)

// Parser интерфейс для реализации парсеров для различных сайтов
type Parser interface {
	Name() string
	Load(ctx context.Context, u url.URL) (BookParser, error)
	CanParse(u url.URL) bool
	AllBooks(ctx context.Context, u url.URL) ([]string, error)
	LoadImage(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error)

	HProxyList(ctx context.Context, u url.URL) (entities.HProxyList, error)
}

type BookParser interface {
	BookDetails(ctx context.Context) (entities.AgentBookDetails, error)
	HProxyBookDetails(ctx context.Context, pageLimit *int) (entities.HProxyBookDetails, error)
}

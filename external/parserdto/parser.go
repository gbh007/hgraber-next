package parserdto

import (
	"context"
	"errors"
	"io"
	"net/url"
)

var (
	InvalidLinkError      = errors.New("invalid link")
	UnknownAttributeError = errors.New("unknown attribute")
)

// Parser интерфейс для реализации парсеров для различных сайтов
type Parser interface {
	Name() string
	CanParse(u url.URL) bool

	Book(ctx context.Context, u url.URL, pageLimit *int) (Book, error)
	Image(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error)
	List(ctx context.Context, u url.URL) (BookList, error)
}

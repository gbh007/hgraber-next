package agentparseruc

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
)

type loader interface {
	HasParser(ctx context.Context, u url.URL) (bool, error)
	Load(ctx context.Context, u url.URL) (parserdto.BookParser, error)
	LoadImage(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error)
	AllBooks(ctx context.Context, u url.URL) ([]string, error)
	HProxyList(ctx context.Context, u url.URL) (parserdto.BookList, error)
}

type UseCase struct {
	logger *slog.Logger

	loader loader
}

func New(
	logger *slog.Logger,
	loader loader,
) *UseCase {
	return &UseCase{
		logger: logger,
		loader: loader,
	}
}

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
	Book(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error)
	Image(ctx context.Context, u url.URL, bookUrl url.URL) (io.ReadCloser, error)
	List(ctx context.Context, u url.URL) (parserdto.BookList, error)
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

package agentcache

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"hgnext/internal/entities"
)

type parseUseCases interface {
	BooksExists(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
	BookByURL(ctx context.Context, u url.URL) (entities.BookContainer, error)
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)
}

type UseCase struct {
	logger *slog.Logger

	parseUseCases parseUseCases
}

func New(
	logger *slog.Logger,
	parseUseCases parseUseCases,
) *UseCase {
	return &UseCase{
		logger:        logger,
		parseUseCases: parseUseCases,
	}
}

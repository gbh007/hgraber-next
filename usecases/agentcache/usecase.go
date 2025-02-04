package agentcache

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
)

type parseUseCases interface {
	BooksExists(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error)
	BookByURL(ctx context.Context, u url.URL) (core.BookContainer, error)
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

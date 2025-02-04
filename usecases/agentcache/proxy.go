package agentcache

import (
	"context"
	"io"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
)

func (uc *UseCase) CheckBooks(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error) {
	return uc.parseUseCases.BooksExists(ctx, urls)
}

func (uc *UseCase) DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error) {
	return uc.parseUseCases.PageBodyByURL(ctx, imageURL)
}

func (uc *UseCase) CheckPages(ctx context.Context, pages []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error) {
	return uc.parseUseCases.PagesExists(ctx, pages)
}

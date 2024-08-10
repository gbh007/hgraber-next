package agentcache

import (
	"context"
	"io"
	"net/url"

	"hgnext/internal/entities"
)

func (uc *UseCase) CheckBooks(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	return uc.parseUseCases.BooksExists(ctx, urls)
}

func (uc *UseCase) DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error) {
	return uc.parseUseCases.PageBodyByURL(ctx, imageURL)
}

func (uc *UseCase) CheckPages(ctx context.Context, pages []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error) {
	return uc.parseUseCases.PagesExists(ctx, pages)
}

package parsing

import (
	"context"

	"hgnext/internal/entities"
)

func (uc *UseCase) BooksToParse(ctx context.Context) ([]entities.Book, error) {
	return uc.storage.UnprocessedBooks(ctx)
}

func (uc *UseCase) PageToLoad(ctx context.Context) ([]entities.PageForDownload, error) {
	return uc.storage.NotDownloadedPages(ctx)
}

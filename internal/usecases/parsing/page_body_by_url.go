package parsing

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"hgnext/internal/entities"
)

func (uc *UseCase) PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error) {
	pages, err := uc.storage.PagesByURL(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("get pages from storage: %w", err)
	}

	for _, p := range pages {
		if !p.IsLoaded() {
			continue
		}

		// TODO: перейти на прямую работу с FS используя page with hash
		body, err := uc.fileStorage.Get(ctx, p.FileID, nil)
		if err != nil {
			return nil, fmt.Errorf("get file from file storage: %w", err)
		}

		return body, nil
	}

	return nil, entities.PageNotFoundError
}

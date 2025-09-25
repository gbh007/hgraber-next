package filesystemusecase

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) PageBody(ctx context.Context, bookID uuid.UUID, pageNumber int) (io.Reader, error) {
	page, err := uc.storage.GetPage(ctx, bookID, pageNumber)
	if err != nil {
		return nil, fmt.Errorf("get page from storage: %w", err)
	}

	if !page.IsLoaded() {
		return nil, fmt.Errorf("missing page body: %w", core.ErrFileNotFound)
	}

	body, err := uc.fileStorage.Get(ctx, page.FileID, nil)
	if err != nil {
		return nil, fmt.Errorf("get file from file storage: %w", err)
	}

	return body, nil
}

package external

import (
	"archive/zip"
	"context"
	"io"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func WriteArchiveAdapter(
	ctx context.Context,
	zipWriter *zip.Writer,
	files interface {
		Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	},
	book entities.BookFull,
) error {
	return WriteArchive(
		ctx,
		zipWriter,
		func(ctx context.Context, pageNumber int) (io.Reader, error) {
			// TODO: перейти на мапу.
			for _, page := range book.Pages {
				if page.PageNumber != pageNumber {
					continue
				}

				// Нет файла, пропускаем
				if page.FileID == uuid.Nil {
					return nil, ErrSkipPageBody
				}

				return files.Get(ctx, page.FileID)
			}

			return nil, ErrSkipPageBody
		},
		Convert(book),
	)
}

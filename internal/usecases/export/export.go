package export

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/external"
	"hgnext/internal/pkg"
)

func (uc *UseCase) Export(ctx context.Context, agentID uuid.UUID, filter entities.BookFilter, deleteAfter bool) error {
	filter.OriginAttributes = true // FIXME: перенести это управление в запрос

	books, err := uc.bookRequester.Books(ctx, filter)
	if err != nil {
		return fmt.Errorf("get books from requester: %w", err)
	}

	uc.tmpStorage.AddToExport(
		pkg.Map(books, func(b entities.BookFull) entities.BookFullWithAgent {
			return entities.BookFullWithAgent{
				BookFull:          b,
				AgentID:           agentID,
				DeleteAfterExport: deleteAfter,
			}
		}),
	)

	return nil
}

func (uc *UseCase) ExportList() []entities.BookFullWithAgent {
	return uc.tmpStorage.ExportList()
}

func (uc *UseCase) ExportArchive(ctx context.Context, book entities.BookFullWithAgent, retry bool) error {
	body, err := uc.newArchive(ctx, book.BookFull)
	if err != nil {
		if retry {
			uc.tmpStorage.AddToExport([]entities.BookFullWithAgent{book})
		}

		return fmt.Errorf("make archive %s: %w", book.Book.ID.String(), err)
	}

	err = uc.agentSystem.ExportArchive(ctx, book.AgentID, entities.AgentExportData{
		BookID:   book.Book.ID,
		BookName: book.Book.Name,
		BookURL:  book.Book.OriginURL,
		Body:     body,
	})
	if err != nil {
		if retry {
			uc.tmpStorage.AddToExport([]entities.BookFullWithAgent{book})
		}

		return fmt.Errorf("export archive %s to agent: %w", book.Book.ID.String(), err)
	}

	if book.DeleteAfterExport {
		err = uc.storage.MarkBookAsDeleted(ctx, book.Book.ID)
		if err != nil {
			if retry {
				uc.tmpStorage.AddToExport([]entities.BookFullWithAgent{book})
			}

			return fmt.Errorf("delete book after export %s: %w", book.Book.ID.String(), err)
		}
	}

	return nil
}

func (uc *UseCase) ExportBook(ctx context.Context, bookID uuid.UUID) (io.Reader, error) {
	book, err := uc.bookRequester.BookOriginFull(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("get book: %w", err)
	}

	return uc.newArchive(ctx, book)
}

func (uc *UseCase) newArchive(ctx context.Context, book entities.BookFull) (io.Reader, error) {
	zipFile := &bytes.Buffer{}
	zipWriter := zip.NewWriter(zipFile)

	err := external.WriteArchiveAdapter(ctx, zipWriter, uc.fileStorage, book)
	if err != nil {
		return nil, fmt.Errorf("write archive: %w", err)
	}

	err = zipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("close archive writer: %w", err)
	}

	return zipFile, nil
}

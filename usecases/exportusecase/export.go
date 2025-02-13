package exportusecase

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/external"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) Export(ctx context.Context, agentID uuid.UUID, filter core.BookFilter, deleteAfter bool) error {
	filter.OriginAttributes = true // FIXME: перенести это управление в запрос

	books, err := uc.storage.BookIDs(ctx, filter)
	if err != nil {
		return fmt.Errorf("get books from requester: %w", err)
	}

	uc.tmpStorage.AddToExport(
		pkg.Map(books, func(bookID uuid.UUID) agentmodel.BookToExport {
			return agentmodel.BookToExport{
				AgentID:           agentID,
				BookID:            bookID,
				DeleteAfterExport: deleteAfter,
			}
		}),
	)

	return nil
}

func (uc *UseCase) ExportList() []agentmodel.BookToExport {
	return uc.tmpStorage.ExportList()
}

func (uc *UseCase) ExportArchive(ctx context.Context, toExport agentmodel.BookToExport, retry bool) error {
	bookContainer, err := uc.bookAdapter.BookRaw(ctx, toExport.BookID)
	if err != nil {
		if retry {
			uc.tmpStorage.AddToExport([]agentmodel.BookToExport{toExport})
		}

		return fmt.Errorf("get book container %s: %w", toExport.BookID.String(), err)
	}

	body, err := uc.newArchive(ctx, bookContainer)
	if err != nil {
		if retry {
			uc.tmpStorage.AddToExport([]agentmodel.BookToExport{toExport})
		}

		return fmt.Errorf("make archive %s: %w", toExport.BookID.String(), err)
	}

	err = uc.agentSystem.ExportArchive(ctx, toExport.AgentID, agentmodel.AgentExportData{
		BookID:   bookContainer.Book.ID,
		BookName: bookContainer.Book.Name,
		BookURL:  bookContainer.Book.OriginURL,
		Body:     body,
	})
	if err != nil {
		if retry {
			uc.tmpStorage.AddToExport([]agentmodel.BookToExport{toExport})
		}

		return fmt.Errorf("export archive %s to agent: %w", toExport.BookID.String(), err)
	}

	if toExport.DeleteAfterExport {
		err = uc.storage.MarkBookAsDeleted(ctx, toExport.BookID)
		if err != nil {
			if retry {
				uc.tmpStorage.AddToExport([]agentmodel.BookToExport{toExport})
			}

			return fmt.Errorf("delete book after export %s: %w", toExport.BookID.String(), err)
		}
	}

	return nil
}

func (uc *UseCase) ExportBook(ctx context.Context, bookID uuid.UUID) (io.Reader, core.BookContainer, error) {
	book, err := uc.bookAdapter.BookRaw(ctx, bookID)
	if err != nil {
		return nil, core.BookContainer{}, fmt.Errorf("get book: %w", err)
	}

	body, err := uc.newArchive(ctx, book)
	if err != nil {
		return nil, core.BookContainer{}, fmt.Errorf("archive book: %w", err)
	}

	return body, book, nil
}

func (uc *UseCase) newArchive(ctx context.Context, book core.BookContainer) (io.Reader, error) {
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

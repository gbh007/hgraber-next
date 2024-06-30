package export

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/external"
	"hgnext/internal/pkg"
)

func (uc *UseCase) Export(ctx context.Context, agentID uuid.UUID, from, to time.Time) error {
	books, err := uc.storage.GetBooks(ctx, entities.BookFilter{
		From: from,
		To:   to,
	})
	if err != nil {
		return fmt.Errorf("get books from storage: %w", err)
	}

	uc.tmpStorage.AddToExport(
		pkg.Map(books, func(b entities.BookFull) entities.BookFullWithAgent {
			return entities.BookFullWithAgent{
				BookFull: b,
				AgentID:  agentID,
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

		return fmt.Errorf("make archive %s: %w", book.ID.String(), err)
	}

	err = uc.agentSystem.ExportArchive(ctx, book.AgentID, book.ID, book.Name, body)
	if err != nil {
		if retry {
			uc.tmpStorage.AddToExport([]entities.BookFullWithAgent{book})
		}

		return fmt.Errorf("export archive %s to agent: %w", book.ID.String(), err)
	}

	return nil
}

func (uc *UseCase) newArchive(ctx context.Context, book entities.BookFull) (io.Reader, error) {
	zipFile := &bytes.Buffer{}
	zipWriter := zip.NewWriter(zipFile)

	w, err := zipWriter.Create("info.json")
	if err != nil {
		return nil, fmt.Errorf("create info: %w", err)
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	err = enc.Encode(external.Convert(book))
	if err != nil {
		return nil, fmt.Errorf("encode info: %w", err)
	}

	for _, p := range book.Pages {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Нет файла, пропускаем
		if p.FileID == uuid.Nil {
			continue
		}

		pageBody, err := uc.fileStorage.Get(ctx, p.FileID)
		if err != nil {
			return nil, fmt.Errorf("get page body: %w", err)
		}

		w, err := zipWriter.Create(fmt.Sprintf("%d%s", p.PageNumber, p.Ext))
		if err != nil {
			return nil, fmt.Errorf("create page %d body: %w", p.PageNumber, err)
		}

		_, err = io.Copy(w, pageBody)
		if err != nil {
			return nil, fmt.Errorf("copy page %d body: %w", p.PageNumber, err)
		}
	}

	err = zipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("close archive writer: %w", err)
	}

	return zipFile, nil
}
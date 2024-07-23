package export

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/external"
)

func (uc *UseCase) ImportArchive(ctx context.Context, body io.Reader) (returnedBookID uuid.UUID, returnedErr error) {
	bodyRaw, err := io.ReadAll(body)
	if err != nil {
		return uuid.Nil, fmt.Errorf("read archive body: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(bodyRaw), int64(len(bodyRaw)))
	if err != nil {
		return uuid.Nil, fmt.Errorf("create zip reader: %w", err)
	}

	book := entities.BookFull{}
	found := false

	for _, f := range zipReader.File {
		if f.Name != "info.json" {
			continue
		}

		found = true

		r, err := f.Open()
		if err != nil {
			return uuid.Nil, fmt.Errorf("open info file: %w", err)
		}

		info := external.Info{}

		err = json.NewDecoder(r).Decode(&info)
		if err != nil {
			return uuid.Nil, fmt.Errorf("decode info file: %w", err)
		}

		err = r.Close()
		if err != nil {
			return uuid.Nil, fmt.Errorf("close info file: %w", err)
		}

		book, err = external.BookToEntity(info.Data)
		if err != nil {
			return uuid.Nil, fmt.Errorf("convert info: %w", err)
		}
	}

	if !found {
		return uuid.Nil, fmt.Errorf("missing book info: %w", entities.BookNotFoundError)
	}

	pageData := make(map[int]uuid.UUID, len(book.Pages))

	for _, f := range zipReader.File {
		number, _ := strconv.Atoi(strings.Split(f.Name, ".")[0])
		if number < 1 {
			continue
		}

		r, err := f.Open()
		if err != nil {
			return uuid.Nil, fmt.Errorf("open page (%d) file: %w", number, err)
		}

		fileID := uuid.Must(uuid.NewV7())

		err = uc.fileStorage.Create(ctx, fileID, r)
		if err != nil {
			return uuid.Nil, fmt.Errorf("page (%d) store file (%s): %w", number, fileID.String(), err)
		}

		err = r.Close()
		if err != nil {
			return uuid.Nil, fmt.Errorf("close page (%d) file: %w", number, err)
		}

		pageData[number] = fileID
	}

	bookID := uuid.Must(uuid.NewV7())
	book.Book.ID = bookID

	if book.Book.CreateAt.IsZero() {
		book.Book.CreateAt = time.Now()
	}

	err = uc.storage.NewBook(ctx, book.Book)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create book: %w", err)
	}

	// Удаляем данные книги если произошла ошибка,
	// данные файлов будет удалять отдельная система синхронизации.
	defer func() {
		if returnedErr == nil {
			return
		}

		uc.logger.Logger(ctx).DebugContext(
			ctx, "try delete book after unsuccess import",
			slog.String("book_id", bookID.String()),
		)

		deleteErr := uc.storage.DeleteBook(ctx, bookID)
		if deleteErr != nil {
			uc.logger.Logger(ctx).ErrorContext(
				ctx, "delete book after unsuccess import",
				slog.Any("error", deleteErr),
				slog.String("book_id", bookID.String()),
			)
		}
	}()

	for _, l := range book.Labels {
		l.BookID = bookID

		err = uc.storage.SetLabel(ctx, l)
		if err != nil {
			return uuid.Nil, fmt.Errorf("set label (%d,%s): %w", l.PageNumber, l.Name, err)
		}
	}

	for code, values := range book.Attributes {
		err = uc.storage.UpdateAttribute(ctx, bookID, code, values)
		if err != nil {
			return uuid.Nil, fmt.Errorf("set attribute (%s): %w", code, err)
		}
	}

	for i, p := range book.Pages {
		fileID, hasFile := pageData[p.PageNumber]
		if hasFile {
			err = uc.storage.NewFile(ctx, entities.File{
				ID:       fileID,
				Filename: fmt.Sprintf("%d%s", p.PageNumber, p.Ext),
				Ext:      p.Ext,
				CreateAt: time.Now(),
			})
			if err != nil {
				return uuid.Nil, fmt.Errorf("page (%d) storage: create file (%s): %w", p.PageNumber, fileID.String(), err)
			}

			p.FileID = fileID
		}

		if p.Downloaded != hasFile {
			uc.logger.Logger(ctx).WarnContext(
				ctx, "mismatch page download status and archive body",
				slog.String("book_id", bookID.String()),
				slog.Int("page_number", p.PageNumber),
				slog.String("file_id", fileID.String()),
				slog.Bool("downloaded", p.Downloaded),
				slog.Bool("has_file", hasFile),
			)
		}

		p.BookID = bookID

		book.Pages[i] = p
	}

	err = uc.storage.UpdateBookPages(ctx, bookID, book.Pages)
	if err != nil {
		return uuid.Nil, fmt.Errorf("set pages: %w", err)
	}

	return bookID, nil
}

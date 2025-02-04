package export

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/external"
)

func (uc *UseCase) ImportArchive(
	ctx context.Context,
	body io.Reader,
	deduplicate bool,
	autoVerify bool,
) (returnedBookID uuid.UUID, returnedErr error) {
	fsID, err := uc.fileStorage.FSIDForDownload(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get fs id for download: %w", err)
	}

	bodyRaw, err := io.ReadAll(body)
	if err != nil {
		return uuid.Nil, fmt.Errorf("read archive body: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(bodyRaw), int64(len(bodyRaw)))
	if err != nil {
		return uuid.Nil, fmt.Errorf("create zip reader: %w", err)
	}

	pageData := make(map[int]uuid.UUID, 50)

	info, err := external.ReadArchive(
		ctx,
		zipReader,
		external.ReadArchiveOptions{
			HandlePageBody: func(ctx context.Context, pageNumber int, _ string, body io.Reader) error {
				fileID := uuid.Must(uuid.NewV7())

				err := uc.fileStorage.Create(ctx, fileID, body, fsID)
				if err != nil {
					return fmt.Errorf("page (%d) store file (%s): %w", pageNumber, fileID.String(), err)
				}

				pageData[pageNumber] = fileID

				return nil
			},
		},
	)

	if errors.Is(err, external.ErrBookInfoNotFound) {
		return uuid.Nil, fmt.Errorf("missing book info: %w", core.BookNotFoundError)
	}

	book, err := external.BookToEntity(info.Data)
	if err != nil {
		return uuid.Nil, fmt.Errorf("convert info: %w", err)
	}

	if deduplicate && book.Book.OriginURL != nil {
		ids, err := uc.storage.GetBookIDsByURL(ctx, *book.Book.OriginURL)
		if err != nil {
			return uuid.Nil, fmt.Errorf("check existing in storage: %w", err)
		}

		// Если есть совпадение, возвращаем первое.
		if len(ids) > 0 {
			return ids[0], nil
		}
	}

	bookID := uuid.Must(uuid.NewV7())
	book.Book.ID = bookID

	if autoVerify {
		book.Book.Verified = true
		book.Book.VerifiedAt = book.Book.CreateAt
	}

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

		uc.logger.DebugContext(
			ctx, "try delete book after unsuccess import",
			slog.String("book_id", bookID.String()),
		)

		deleteErr := uc.storage.DeleteBook(ctx, bookID)
		if deleteErr != nil {
			uc.logger.ErrorContext(
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

	err = uc.storage.UpdateOriginAttributes(ctx, bookID, book.Attributes)
	if err != nil {
		return uuid.Nil, fmt.Errorf("set original attributes: %w", err)
	}

	err = uc.storage.UpdateAttributes(ctx, bookID, book.Attributes)
	if err != nil {
		return uuid.Nil, fmt.Errorf("set attributes: %w", err)
	}

	for i, p := range book.Pages {
		fileID, hasFile := pageData[p.PageNumber]
		if hasFile {
			err = uc.storage.NewFile(ctx, core.File{
				ID:       fileID,
				Filename: fmt.Sprintf("%d%s", p.PageNumber, p.Ext),
				Ext:      p.Ext,
				FSID:     fsID,
				CreateAt: time.Now(),
			})
			if err != nil {
				return uuid.Nil, fmt.Errorf("page (%d) storage: create file (%s): %w", p.PageNumber, fileID.String(), err)
			}

			p.FileID = fileID
		}

		if p.Downloaded != hasFile {
			uc.logger.WarnContext(
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

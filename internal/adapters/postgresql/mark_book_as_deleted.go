package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/internal/entities"
)

func (d *Database) MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			d.logger.ErrorContext(
				ctx, "rollback MarkBookAsDeleted tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.ExecContext(ctx, `INSERT INTO
    deleted_pages
SELECT
    p.book_id,
    p.page_number,
    p.ext,
    p.origin_url,
    f.md5_sum,
    f.sha256_sum,
    f.size,
    p.downloaded,
    p.create_at AS created_at,
    p.load_at AS loaded_at
FROM pages p
    LEFT JOIN files f ON p.file_id = f.id
WHERE
    p.book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("copy pages: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM pages WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("delete pages: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("delete attributes: %w", err)
	}

	res, err := tx.ExecContext(
		ctx,
		`UPDATE books SET deleted_at = $2, deleted = $3 WHERE id = $1;`,
		bookID.String(), time.Now().UTC(), true,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if !d.isApply(ctx, res) {
		return entities.BookNotFoundError
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *Database) MarkPageAsDeleted(ctx context.Context, bookID uuid.UUID, pageNumber int) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			d.logger.ErrorContext(
				ctx, "rollback MarkPageAsDeleted tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.ExecContext(ctx, `INSERT INTO
    deleted_pages
SELECT
    p.book_id,
    p.page_number,
    p.ext,
    p.origin_url,
    f.md5_sum,
    f.sha256_sum,
    f.size,
    p.downloaded,
    p.create_at AS created_at,
    p.load_at AS loaded_at
FROM pages p
    LEFT JOIN files f ON p.file_id = f.id
WHERE
    p.book_id = $1 AND p.page_number = $2;`, bookID.String(), pageNumber)
	if err != nil {
		return fmt.Errorf("copy page: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM pages WHERE book_id = $1 AND page_number = $2;`, bookID.String(), pageNumber)
	if err != nil {
		return fmt.Errorf("delete page: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *Database) UpdateBookDeletion(ctx context.Context, book entities.Book) error {
	builder := squirrel.Update("books").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(
			map[string]interface{}{
				"deleted":    book.Deleted,
				"deleted_at": model.TimeToDB(book.DeletedAt),
			},
		).
		Where(squirrel.Eq{
			"id": book.ID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("storage: build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("storage: exec query: %w", err)
	}

	if !d.isApply(ctx, res) {
		return entities.BookNotFoundError
	}

	return nil
}

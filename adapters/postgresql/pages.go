package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (core.Page, error) {
	builder := squirrel.Select(model.PageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"book_id":     id,
			"page_number": pageNumber,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.Page{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	page := core.Page{}

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.PageScanner(&page))

	if err != nil {
		return core.Page{}, fmt.Errorf("exec query :%w", err)
	}

	return page, nil
}

func (d *Database) UpdatePageDownloaded(ctx context.Context, id uuid.UUID, pageNumber int, downloaded bool, fileID uuid.UUID) error {
	res, err := d.pool.Exec(
		ctx,
		`UPDATE pages SET downloaded = $1, load_at = $2, file_id = $5 WHERE book_id = $3 AND page_number = $4;`,
		downloaded, time.Now().UTC(), id, pageNumber, model.UUIDToDB(fileID),
	)
	if err != nil {
		return err
	}

	if res.RowsAffected() < 1 {
		return core.PageNotFoundError
	}

	return nil
}

// FIXME: отрефакторить на squirel
func (d *Database) UpdateBookPages(ctx context.Context, id uuid.UUID, pages []core.Page) error {
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateBookPages tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM pages WHERE book_id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete old pages: %w", err)
	}

	// TODO: слить с аналогичным дейтвием, реализовать как приватную функцию которая принимает транзакцию.
	for _, v := range pages {
		_, err = tx.Exec(
			ctx,
			`INSERT INTO pages (book_id, page_number, ext, origin_url, create_at, downloaded, load_at, file_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
			id, v.PageNumber, v.Ext, model.URLToDB(v.OriginURL), v.CreateAt.UTC(), v.Downloaded, model.TimeToDB(v.LoadAt), model.UUIDToDB(v.FileID),
		)
		if err != nil {
			return fmt.Errorf("insert page %d: %w", v.PageNumber, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// FIXME: отрефакторить на squirel
func (d *Database) NewBookPages(ctx context.Context, pages []core.Page) error {
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateBookPages tx",
				slog.Any("err", err),
			)
		}
	}()

	// TODO: слить с аналогичным дейтвием, реализовать как приватную функцию которая принимает транзакцию.
	for _, v := range pages {
		_, err = tx.Exec(
			ctx,
			`INSERT INTO pages (book_id, page_number, ext, origin_url, create_at, downloaded, load_at, file_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
			v.BookID, v.PageNumber, v.Ext, model.URLToDB(v.OriginURL), v.CreateAt.UTC(), v.Downloaded, model.TimeToDB(v.LoadAt), model.UUIDToDB(v.FileID),
		)
		if err != nil {
			return fmt.Errorf("insert page %d: %w", v.PageNumber, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *Database) BookPages(ctx context.Context, bookID uuid.UUID) ([]core.Page, error) {
	builder := squirrel.Select(model.PageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"book_id": bookID,
		}).
		OrderBy("page_number")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]core.Page, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.Page{}

		err := rows.Scan(model.PageScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, page)
	}

	return result, nil
}

func (d *Database) PagesByURL(ctx context.Context, u url.URL) ([]core.Page, error) {
	builder := squirrel.Select(model.PageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"origin_url": u.String(),
		}).
		OrderBy("book_id", "page_number")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]core.Page, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.Page{}

		err := rows.Scan(model.PageScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, page)
	}

	return result, nil
}

func (d *Database) BookPagesCount(ctx context.Context, bookID uuid.UUID) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"book_id": bookID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	count := sql.NullInt64{}
	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get count :%w", err)
	}

	return int(count.Int64), nil
}

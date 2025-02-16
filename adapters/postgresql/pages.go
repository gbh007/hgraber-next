package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (core.Page, error) {
	raw := new(model.Page)

	err := d.db.GetContext(
		ctx, raw,
		`SELECT * FROM pages WHERE book_id = $1 AND page_number = $2 LIMIT 1;`,
		id, pageNumber,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return core.Page{}, core.PageNotFoundError
	}

	if err != nil {
		return core.Page{}, fmt.Errorf("get page from db: %w", err)
	}

	p, err := raw.ToEntity()
	if err != nil {
		return core.Page{}, fmt.Errorf("convert page: %w", err)
	}

	return p, nil
}

func (d *Database) UpdatePageDownloaded(ctx context.Context, id uuid.UUID, pageNumber int, downloaded bool, fileID uuid.UUID) error {
	res, err := d.db.ExecContext(
		ctx,
		`UPDATE pages SET downloaded = $1, load_at = $2, file_id = $5 WHERE book_id = $3 AND page_number = $4;`,
		downloaded, time.Now().UTC(), id, pageNumber, model.UUIDToDB(fileID),
	)
	if err != nil {
		return err
	}

	if !d.isApply(ctx, res) {
		return core.PageNotFoundError
	}

	return nil
}

// FIXME: отрефакторить на squirel
func (d *Database) UpdateBookPages(ctx context.Context, id uuid.UUID, pages []core.Page) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM pages WHERE book_id = $1;`, id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			d.logger.ErrorContext(ctx, rollbackErr.Error())
		}

		return err
	}

	for _, v := range pages {
		_, err = tx.ExecContext(
			ctx,
			`INSERT INTO pages (book_id, page_number, ext, origin_url, create_at, downloaded, load_at, file_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
			id, v.PageNumber, v.Ext, model.URLToDB(v.OriginURL), v.CreateAt.UTC(), v.Downloaded, model.TimeToDB(v.LoadAt), model.UUIDToDB(v.FileID),
		)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				d.logger.ErrorContext(ctx, rollbackErr.Error())
			}

			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// FIXME: отрефакторить на squirel
func (d *Database) NewBookPages(ctx context.Context, pages []core.Page) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	for _, v := range pages {
		_, err = tx.ExecContext(
			ctx,
			`INSERT INTO pages (book_id, page_number, ext, origin_url, create_at, downloaded, load_at, file_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
			v.BookID, v.PageNumber, v.Ext, model.URLToDB(v.OriginURL), v.CreateAt.UTC(), v.Downloaded, model.TimeToDB(v.LoadAt), model.UUIDToDB(v.FileID),
		)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				d.logger.ErrorContext(ctx, rollbackErr.Error())
			}

			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) bookPages(ctx context.Context, bookID uuid.UUID) ([]*model.Page, error) {
	raw := make([]*model.Page, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM pages WHERE book_id = $1 ORDER BY page_number;`, bookID)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (d *Database) BookPages(ctx context.Context, bookID uuid.UUID) ([]core.Page, error) {
	pages, err := d.bookPages(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("get pages :%w", err)
	}

	out := make([]core.Page, 0, len(pages))

	for _, pageRaw := range pages {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert page :%w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) PagesByURL(ctx context.Context, u url.URL) ([]core.Page, error) {
	raw := make([]*model.Page, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM pages WHERE origin_url = $1 ORDER BY book_id, page_number;`, u.String())
	if err != nil {
		return nil, fmt.Errorf("get pages :%w", err)
	}

	out := make([]core.Page, 0, len(raw))

	for _, pageRaw := range raw {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert page :%w", err)
		}

		out = append(out, page)
	}

	return out, nil
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

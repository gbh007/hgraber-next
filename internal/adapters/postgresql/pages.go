package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
)

func (d *Database) GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (entities.Page, error) {
	raw := new(model.Page)

	err := d.db.GetContext(
		ctx, raw,
		`SELECT * FROM pages WHERE book_id = $1 AND page_number = $2 LIMIT 1;`,
		id.String(), pageNumber,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return entities.Page{}, entities.PageNotFoundError
	}

	if err != nil {
		return entities.Page{}, fmt.Errorf("get page from db: %w", err)
	}

	p, err := raw.ToEntity()
	if err != nil {
		return entities.Page{}, fmt.Errorf("convert page: %w", err)
	}

	return p, nil
}

func (d *Database) UpdatePageDownloaded(ctx context.Context, id uuid.UUID, pageNumber int, downloaded bool, fileID uuid.UUID) error {
	res, err := d.db.ExecContext(
		ctx,
		`UPDATE pages SET downloaded = $1, load_at = $2, file_id = $5 WHERE book_id = $3 AND page_number = $4;`,
		downloaded, time.Now().UTC(), id.String(), pageNumber, model.UUIDToDB(fileID),
	)
	if err != nil {
		return err
	}

	if !d.isApply(ctx, res) {
		return entities.PageNotFoundError
	}

	// Состояние размера изменилось, сбрасываем кеши.
	d.cachePageFileSize.Store(0)
	d.cacheFileSize.Store(0)
	d.cacheDownloadedBookCount.Store(0)
	d.cacheVerifiedBookCount.Store(0)

	return nil
}

// FIXME: отрефакторить на squirel
func (d *Database) UpdateBookPages(ctx context.Context, id uuid.UUID, pages []entities.Page) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM pages WHERE book_id = $1;`, id.String())
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
			id.String(), v.PageNumber, v.Ext, model.URLToDB(v.OriginURL), v.CreateAt.UTC(), v.Downloaded, model.TimeToDB(v.LoadAt), model.UUIDToDB(v.FileID),
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

	// Состояние размера изменилось, сбрасываем кеши.
	d.cachePageFileSize.Store(0)
	d.cacheFileSize.Store(0)
	d.cachePageCount.Store(0)
	d.cacheDownloadedBookCount.Store(0)
	d.cacheVerifiedBookCount.Store(0)

	return nil
}

// FIXME: отрефакторить на squirel
func (d *Database) NewBookPages(ctx context.Context, pages []entities.Page) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	for _, v := range pages {
		_, err = tx.ExecContext(
			ctx,
			`INSERT INTO pages (book_id, page_number, ext, origin_url, create_at, downloaded, load_at, file_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
			v.BookID.String(), v.PageNumber, v.Ext, model.URLToDB(v.OriginURL), v.CreateAt.UTC(), v.Downloaded, model.TimeToDB(v.LoadAt), model.UUIDToDB(v.FileID),
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

	// Состояние размера изменилось, сбрасываем кеши.
	d.cachePageFileSize.Store(0)
	d.cacheFileSize.Store(0)
	d.cachePageCount.Store(0)
	d.cacheDownloadedBookCount.Store(0)
	d.cacheVerifiedBookCount.Store(0)

	return nil
}

func (d *Database) bookPages(ctx context.Context, bookID uuid.UUID) ([]*model.Page, error) {
	raw := make([]*model.Page, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM pages WHERE book_id = $1 ORDER BY page_number;`, bookID.String())
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (d *Database) BookPages(ctx context.Context, bookID uuid.UUID) ([]entities.Page, error) {
	pages, err := d.bookPages(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("get pages :%w", err)
	}

	out := make([]entities.Page, 0, len(pages))

	for _, pageRaw := range pages {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert page :%w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) PagesByURL(ctx context.Context, u url.URL) ([]entities.Page, error) {
	raw := make([]*model.Page, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM pages WHERE origin_url = $1 ORDER BY book_id, page_number;`, u.String())
	if err != nil {
		return nil, fmt.Errorf("get pages :%w", err)
	}

	out := make([]entities.Page, 0, len(raw))

	for _, pageRaw := range raw {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert page :%w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

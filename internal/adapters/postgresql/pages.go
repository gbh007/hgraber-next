package postgresql

import (
	"context"
	"database/sql"
	"errors"
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
		return entities.Page{}, err
	}

	p, err := raw.ToEntity()
	if err != nil {
		return entities.Page{}, err
	}

	return p, nil
}

func (d *Database) GetNotDownloadedPages(ctx context.Context) []entities.Page {
	raw := make([]*model.Page, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM pages WHERE downloaded = FALSE;`)
	if err != nil {
		d.logger.ErrorContext(ctx, err.Error())

		return []entities.Page{}
	}

	out := make([]entities.Page, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			d.logger.ErrorContext(ctx, err.Error())

			return []entities.Page{}
		}
	}

	return out
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

	return nil
}

func (d *Database) UpdatePage(ctx context.Context, id uuid.UUID, pageNumber int, downloaded bool, url string) error {
	res, err := d.db.ExecContext(
		ctx,
		`UPDATE pages SET downloaded = $1, load_at = $2, url = $5 WHERE book_id = $3 AND page_number = $4;`,
		downloaded, time.Now().UTC(), id.String(), pageNumber, url,
	)
	if err != nil {
		return err
	}

	if !d.isApply(ctx, res) {
		return entities.PageNotFoundError
	}

	return nil
}

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

	return nil
}

func (d *Database) getBookPages(ctx context.Context, bookID uuid.UUID) ([]*model.Page, error) {
	raw := make([]*model.Page, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM pages WHERE book_id = $1 ORDER BY page_number;`, bookID.String())
	if err != nil {
		return nil, err
	}

	return raw, nil
}

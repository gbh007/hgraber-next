package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) NewBook(ctx context.Context, book core.Book) error {
	builder := squirrel.Insert("books").
		PlaceholderFormat(squirrel.Dollar).SetMap(
		map[string]interface{}{
			"id":                book.ID,
			"name":              model.StringToDB(book.Name),
			"origin_url":        model.URLToDB(book.OriginURL),
			"page_count":        model.Int32ToDB(book.PageCount),
			"attributes_parsed": book.AttributesParsed,
			"verified":          book.Verified,
			"verified_at":       model.TimeToDB(book.VerifiedAt),
			"is_rebuild":        book.IsRebuild,
			"create_at":         book.CreateAt,
		},
	)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("storage: build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("storage: exec query: %w", err)
	}

	return nil
}

func (d *Database) UpdateBook(ctx context.Context, book core.Book) error {
	builder := squirrel.Update("books").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(
			map[string]interface{}{
				"name":              model.StringToDB(book.Name),
				"origin_url":        model.URLToDB(book.OriginURL),
				"page_count":        model.Int32ToDB(book.PageCount),
				"attributes_parsed": book.AttributesParsed,
				"verified":          book.Verified,
				"verified_at":       model.TimeToDB(book.VerifiedAt),
				"is_rebuild":        book.IsRebuild,
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
		return core.BookNotFoundError
	}

	return nil
}

func (d *Database) GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error) {
	var idsRaw []string

	err := d.db.SelectContext(ctx, &idsRaw, `SELECT id FROM books WHERE origin_url = $1;`, url.String())
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(idsRaw))

	for i, idRaw := range idsRaw {
		ids[i], err = uuid.Parse(idRaw)
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

func (d *Database) GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error) {
	raw := new(model.Book)

	err := d.db.GetContext(ctx, raw, `SELECT * FROM books WHERE id = $1 LIMIT 1;`, bookID)
	if errors.Is(err, sql.ErrNoRows) {
		return core.Book{}, fmt.Errorf("%w - %d", core.BookNotFoundError, bookID)
	}

	if err != nil {
		return core.Book{}, err
	}

	b, err := raw.ToEntity()
	if err != nil {
		return core.Book{}, err
	}

	return b, nil
}

func (d *Database) DeleteBook(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("books").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if !d.isApply(ctx, res) {
		return core.BookNotFoundError
	}

	return nil
}

func (d *Database) DeleteBooks(ctx context.Context, ids []uuid.UUID) error {
	builder := squirrel.Delete("books").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": ids,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if !d.isApply(ctx, res) {
		return core.BookNotFoundError
	}

	return nil
}

func (d *Database) BookIDsWithDeletedRebuilds(ctx context.Context) ([]uuid.UUID, error) {
	builder := squirrel.Select("id").
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.Eq{
			"deleted":    true,
			"is_rebuild": true,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := []uuid.UUID{}

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		id := uuid.UUID{}

		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, id)
	}

	return result, nil
}

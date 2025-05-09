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

	_, err = d.pool.Exec(ctx, query, args...)
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

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("storage: exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.BookNotFoundError
	}

	return nil
}

func (d *Database) GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error) {
	builder := squirrel.Select("id").
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.Eq{
			"origin_url": url.String(),
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

func (d *Database) GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error) {
	builder := squirrel.Select(model.BookColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.Eq{
			"id": bookID,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.Book{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	book := core.Book{}

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.BookScanner(&book))

	if errors.Is(err, sql.ErrNoRows) {
		return core.Book{}, core.BookNotFoundError
	}

	if err != nil {
		return core.Book{}, fmt.Errorf("exec query :%w", err)
	}

	return book, nil
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

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
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

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
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

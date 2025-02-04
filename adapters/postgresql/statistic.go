package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/entities"
)

func (d *Database) BooksCountByAuthor(ctx context.Context) (map[string]int64, error) {
	builder := squirrel.Select("COUNT(*)", "value").
		PlaceholderFormat(squirrel.Dollar).
		From("book_attributes").
		Where(squirrel.Eq{
			"attr": "author",
		}).
		GroupBy("value")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	out := make(map[string]int64, 100)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count sql.NullInt64
			name  string
		)

		err = rows.Scan(&count, &name)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[name] = count.Int64
	}

	return out, nil
}

func (d *Database) PageSizeByAuthor(ctx context.Context) (map[string]entities.SizeWithCount, error) {
	builder := squirrel.Select("COUNT(*)", "ba.value", "SUM(f.size)").
		PlaceholderFormat(squirrel.Dollar).
		From("book_attributes ba").
		InnerJoin("pages p ON ba.book_id = p.book_id").
		InnerJoin("files f ON f.id = p.file_id").
		Where(squirrel.Eq{
			"ba.attr": "author",
		}).
		GroupBy("ba.value")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	out := make(map[string]entities.SizeWithCount, 100)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count sql.NullInt64
			size  sql.NullInt64
			name  string
		)

		err = rows.Scan(&count, &name, &size)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[name] = entities.SizeWithCount{
			Count: count.Int64,
			Size:  size.Int64,
		}
	}

	return out, nil
}

func (d *Database) BookSizes(ctx context.Context) (map[uuid.UUID]entities.SizeWithCount, error) {
	builder := squirrel.Select("COUNT(*)", "p.book_id", "SUM(f.size)").
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		InnerJoin("files f ON f.id = p.file_id").
		GroupBy("p.book_id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	out := make(map[uuid.UUID]entities.SizeWithCount, 100)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count  sql.NullInt64
			size   sql.NullInt64
			bookID uuid.UUID
		)

		err = rows.Scan(&count, &bookID, &size)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[bookID] = entities.SizeWithCount{
			Count: count.Int64,
			Size:  size.Int64,
		}
	}

	return out, nil
}

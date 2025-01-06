package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
)

func (d *Database) DeletedPagesHashes(ctx context.Context) ([]entities.FileHash, error) {
	builder := squirrel.Select(
		"md5_sum",
		"sha256_sum",
		"size",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("deleted_pages").
		Where(squirrel.And{
			squirrel.Expr(`md5_sum IS NOT NULL`),
			squirrel.Expr(`sha256_sum IS NOT NULL`),
			squirrel.Expr(`size IS NOT NULL`),
		}).
		GroupBy(
			"md5_sum",
			"sha256_sum",
			"size",
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]entities.FileHash, 0, 100)

	for rows.Next() {
		hash := entities.FileHash{}

		err = rows.Scan(
			&hash.Md5Sum,
			&hash.Sha256Sum,
			&hash.Size,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, hash)
	}

	return result, nil
}

func (d *Database) DeletedPages(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error) {
	builder := squirrel.Select(model.DeletedPageToPageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("deleted_pages").
		Where(squirrel.Eq{
			"book_id": bookID,
		}).
		OrderBy("page_number")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	out := make([]entities.PageWithHash, 0, 10)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := entities.PageWithHash{}

		err := rows.Scan(model.DeletedPageToPageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) RemoveDeletedPages(ctx context.Context, bookID uuid.UUID, pageNumbers []int) error {
	builder := squirrel.Delete("deleted_pages").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id":     bookID,
			"page_number": pageNumbers,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query :%w", err)
	}

	return nil
}

func (d *Database) TruncateDeletedPages(ctx context.Context) error {
	_, err := d.pool.Exec(ctx, `TRUNCATE deleted_pages;`)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

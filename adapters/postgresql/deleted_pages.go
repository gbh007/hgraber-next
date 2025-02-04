package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) DeletedPagesHashes(ctx context.Context) ([]core.FileHash, error) {
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

	result := make([]core.FileHash, 0, 100)

	for rows.Next() {
		hash := core.FileHash{}

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

func (d *Database) DeletedPages(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error) {
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

	out := make([]core.PageWithHash, 0, 10)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.PageWithHash{}

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

func (d *Database) RemoveDeletedPagesByHash(ctx context.Context, hash core.FileHash) error {
	builder := squirrel.Delete("deleted_pages").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"md5_sum":    hash.Md5Sum,
			"sha256_sum": hash.Sha256Sum,
			"\"size\"":   hash.Size,
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

func (d *Database) RemoveDeletedPagesByHashes(ctx context.Context, hashes []core.FileHash) error {
	batch := &pgx.Batch{}

	resultCount := 0

	for _, hash := range hashes {
		builder := squirrel.Delete("deleted_pages").
			PlaceholderFormat(squirrel.Dollar).
			Where(squirrel.Eq{
				"md5_sum":    hash.Md5Sum,
				"sha256_sum": hash.Sha256Sum,
				"\"size\"":   hash.Size,
			})

		query, args, err := builder.ToSql()
		if err != nil {
			return fmt.Errorf("build query: %w", err)
		}

		d.squirrelDebugLog(ctx, query, args)
		batch.Queue(query, args...)

		resultCount++
	}

	batchResult := d.pool.SendBatch(ctx, batch)
	defer batchResult.Close()

	for range resultCount {
		_, err := batchResult.Exec()
		if err != nil {
			return fmt.Errorf("exec query: %w", err)
		}
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

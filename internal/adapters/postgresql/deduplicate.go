package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
)

func (d *Database) BookIDsByMD5(ctx context.Context, md5Sums []string) ([]uuid.UUID, error) {
	builder := squirrel.Select("b.id").
		PlaceholderFormat(squirrel.Dollar).
		From("books b").
		InnerJoin("pages p ON p.book_id = b.id").
		InnerJoin("files f ON f.id = p.file_id").
		Where(squirrel.Eq{
			"f.md5_sum": md5Sums,
		}).
		GroupBy("b.id")

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

func (d *Database) BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error) {
	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"p.book_id": bookID,
		}).
		OrderBy("p.page_number")

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

		err := rows.Scan(model.PageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) BookPageWithHash(ctx context.Context, bookID uuid.UUID, pageNumber int) (entities.PageWithHash, error) {
	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"p.book_id":     bookID,
			"p.page_number": pageNumber,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return entities.PageWithHash{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	page := entities.PageWithHash{}
	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.PageWithHashScanner(&page))
	if err != nil {
		return entities.PageWithHash{}, fmt.Errorf("exec query :%w", err)
	}

	return page, nil
}

func (d *Database) BookPagesWithHashByHash(ctx context.Context, hash entities.FileHash) ([]entities.PageWithHash, error) {
	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"f.md5_sum":    hash.Md5Sum,
			"f.sha256_sum": hash.Sha256Sum,
			"f.size":       hash.Size,
		})

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

		err := rows.Scan(model.PageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) BookPagesWithHashByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.PageWithHash, error) {
	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"f.md5_sum": md5Sums,
		})

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

		err := rows.Scan(model.PageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) BookPagesCountByHash(ctx context.Context, hash entities.FileHash) (int64, error) {
	builder := squirrel.Select("COUNT(*)").
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"f.md5_sum":    hash.Md5Sum,
			"f.sha256_sum": hash.Sha256Sum,
			"f.size":       hash.Size,
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

	return count.Int64, nil
}

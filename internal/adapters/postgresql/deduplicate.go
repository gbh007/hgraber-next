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

func (d *Database) BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error) {
	result := []uuid.UUID{}

	err := d.db.SelectContext(ctx, &result, `SELECT b.id FROM books b
INNER JOIN pages p ON p.book_id = b.id
INNER JOIN files f ON f.id = p.file_id 
WHERE f.md5_sum = ANY ($1)
GROUP BY b.id;`, md5sums)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Database) BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error) {
	pages := make([]model.PageWithHash, 0)

	err := d.db.SelectContext(ctx, &pages, `SELECT p.book_id, p.page_number, p.ext, p.origin_url, p.downloaded, p.file_id, f.md5_sum, f.sha256_sum, f."size"
FROM pages p left join files f on p.file_id = f.id
WHERE p.book_id = $1 ORDER BY page_number;`, bookID)
	if err != nil {
		return nil, fmt.Errorf("get pages :%w", err)
	}

	out := make([]entities.PageWithHash, 0, len(pages))

	for _, pageRaw := range pages {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert page :%w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

func (d *Database) BookPageWithHash(ctx context.Context, bookID uuid.UUID, pageNumber int) (entities.PageWithHash, error) {
	pageRaw := model.PageWithHash{}

	err := d.db.GetContext(ctx, &pageRaw, `SELECT p.book_id, p.page_number, p.ext, p.origin_url, p.downloaded, p.file_id, f.md5_sum, f.sha256_sum, f."size"
FROM pages p left join files f on p.file_id = f.id
WHERE p.book_id = $1 AND page_number = $2;`, bookID, pageNumber)
	if err != nil {
		return entities.PageWithHash{}, fmt.Errorf("get page :%w", err)
	}

	page, err := pageRaw.ToEntity()
	if err != nil {
		return entities.PageWithHash{}, fmt.Errorf("convert page :%w", err)
	}

	return page, nil
}

func (d *Database) BookPagesWithHashByHash(ctx context.Context, hash entities.FileHash) ([]entities.PageWithHash, error) {
	builder := squirrel.Select(
		"p.book_id",
		"p.page_number",
		"p.ext",
		"p.origin_url",
		"p.downloaded",
		"p.file_id",
		"f.md5_sum",
		"f.sha256_sum",
		"f.size",
	).
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

	pages := make([]model.PageWithHash, 0)

	err = d.db.SelectContext(ctx, &pages, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get pages :%w", err)
	}

	out := make([]entities.PageWithHash, 0, len(pages))

	for _, pageRaw := range pages {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert page :%w", err)
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

	raw := sql.NullInt64{}
	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(&raw)
	if err != nil {
		return 0, fmt.Errorf("get count :%w", err)
	}

	return raw.Int64, nil
}

package postgresql

import (
	"context"
	"fmt"

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

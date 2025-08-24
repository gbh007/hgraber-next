package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPagesWithHashByHash(ctx context.Context, hash core.FileHash) ([]core.PageWithHash, error) {
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

	repo.SquirrelDebugLog(ctx, query, args)

	out := make([]core.PageWithHash, 0, 10) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.PageWithHash{}

		err := rows.Scan(model.PageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

package page

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPagesCountByHash(ctx context.Context, hash core.FileHash) (int64, error) {
	pageTable := model.PageTable

	builder := squirrel.Select("COUNT(*)").
		PlaceholderFormat(squirrel.Dollar).
		From(pageTable.Name() + " p").
		LeftJoin("files f ON p." + pageTable.ColumnFileID() + " = f.id").
		Where(squirrel.Eq{
			"f.md5_sum":    hash.Md5Sum,
			"f.sha256_sum": hash.Sha256Sum,
			"f.size":       hash.Size,
		})

	query, args := builder.MustSql()

	count := sql.NullInt64{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get count: %w", err)
	}

	return count.Int64, nil
}

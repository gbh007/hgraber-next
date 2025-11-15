package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) DeletedPagesHashes(ctx context.Context) ([]core.FileHash, error) {
	deletedPageTable := model.DeletedPageTable

	builder := squirrel.Select(
		deletedPageTable.ColumnMd5Sum(),
		deletedPageTable.ColumnSha256Sum(),
		deletedPageTable.ColumnSize(),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(deletedPageTable.Name()).
		Where(squirrel.And{
			squirrel.Expr(deletedPageTable.ColumnMd5Sum() + " IS NOT NULL"),
			squirrel.Expr(deletedPageTable.ColumnSha256Sum() + " IS NOT NULL"),
			squirrel.Expr(deletedPageTable.ColumnSize() + " IS NOT NULL"),
		}).
		GroupBy(
			deletedPageTable.ColumnMd5Sum(),
			deletedPageTable.ColumnSha256Sum(),
			deletedPageTable.ColumnSize(),
		)

	query, args := builder.MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.FileHash, 0, 100) //nolint:mnd // оптимизация

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

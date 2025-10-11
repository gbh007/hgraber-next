package deadhash

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *DeadHashRepo) SetDeadHash(ctx context.Context, hash core.DeadHash) error {
	table := model.DeadHashTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnMd5Sum():    hash.Md5Sum,
			table.ColumnSha256Sum(): hash.Sha256Sum,
			table.ColumnSize():      hash.Size,
			table.ColumnCreatedAt(): hash.CreatedAt,
		}).
		Suffix(`ON CONFLICT DO NOTHING`)

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

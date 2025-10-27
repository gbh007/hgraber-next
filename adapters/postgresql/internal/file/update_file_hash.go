package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) UpdateFileHash(ctx context.Context, id uuid.UUID, md5Sum, sha256Sum string, size int64) error {
	table := model.FileTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnMd5Sum():    model.StringToDB(md5Sum),
			table.ColumnSha256Sum(): model.StringToDB(sha256Sum),
			table.ColumnSize():      model.Int64ToDB(size),
		}).
		Where(squirrel.Eq{
			table.ColumnID(): id,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrFileNotFound
	}

	return nil
}

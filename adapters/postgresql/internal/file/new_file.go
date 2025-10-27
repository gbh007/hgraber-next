package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) NewFile(ctx context.Context, file core.File) error {
	table := model.FileTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnID():       file.ID,
			table.ColumnFilename(): file.Filename,
			table.ColumnExt():      file.Ext,
			table.ColumnCreateAt(): file.CreateAt,
			table.ColumnFSID():     file.FSID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

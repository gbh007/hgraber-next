package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) NewFile(ctx context.Context, file core.File) error {
	builder := squirrel.Insert("files").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"id":        file.ID,
			"filename":  file.Filename,
			"ext":       file.Ext,
			"create_at": file.CreateAt,
			"fs_id":     file.FSID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

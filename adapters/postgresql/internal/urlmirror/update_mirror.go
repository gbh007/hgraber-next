package urlmirror

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (repo *URLMirrorRepo) UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error {
	table := model.URLMirrorTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnName():        mirror.Name,
			table.ColumnPrefixes():    mirror.Prefixes,
			table.ColumnDescription(): model.StringToDB(mirror.Description),
		}).
		Where(squirrel.Eq{
			table.ColumnID(): mirror.ID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

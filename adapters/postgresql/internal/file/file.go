package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) File(ctx context.Context, id uuid.UUID) (core.File, error) {
	table := model.FileTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnID(): id,
		}).
		Limit(1)

	query, args := builder.MustSql()

	file := core.File{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(table.Scanner(&file))
	if err != nil {
		return core.File{}, fmt.Errorf("exec: %w", err)
	}

	return file, nil
}

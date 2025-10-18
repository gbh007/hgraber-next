package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (repo *FileRepo) FileStorage(ctx context.Context, id uuid.UUID) (fsmodel.FileStorageSystem, error) {
	table := model.FileStorageTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnID(): id,
		}).
		Limit(1)

	query, args := builder.MustSql()

	fs := fsmodel.FileStorageSystem{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(table.Scanner(&fs))
	if err != nil {
		return fsmodel.FileStorageSystem{}, fmt.Errorf("scan: %w", err)
	}

	return fs, nil
}

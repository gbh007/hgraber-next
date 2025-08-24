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
	builder := squirrel.Select(model.FileStorageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("file_storages").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return fsmodel.FileStorageSystem{}, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	fs := fsmodel.FileStorageSystem{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.FileStorageScanner(&fs))
	if err != nil {
		return fsmodel.FileStorageSystem{}, fmt.Errorf("scan: %w", err)
	}

	return fs, nil
}

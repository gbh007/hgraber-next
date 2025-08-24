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
	builder := squirrel.Select(model.FileColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.File{}, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	file := core.File{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.FileScanner(&file))
	if err != nil {
		return core.File{}, fmt.Errorf("exec: %w", err)
	}

	return file, nil
}

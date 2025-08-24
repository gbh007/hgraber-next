package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *FileRepo) DeleteFileStorage(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("file_storages").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *PageRepo) ReplaceFile(ctx context.Context, oldFileID, newFileID uuid.UUID) error {
	builder := squirrel.Update("pages").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"file_id": newFileID,
		}).
		Where(squirrel.Eq{
			"file_id": oldFileID,
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

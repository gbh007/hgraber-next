package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *PageRepo) ReplaceFile(ctx context.Context, oldFileID, newFileID uuid.UUID) error {
	table := model.PageTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnFileID(): newFileID,
		}).
		Where(squirrel.Eq{
			table.ColumnFileID(): oldFileID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

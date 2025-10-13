package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error) {
	table := model.BookLabelTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnBookID(): bookID,
		})

	query, args := builder.MustSql()

	result := make([]core.BookLabel, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		label := core.BookLabel{}

		err := rows.Scan(table.Scanner(&label))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, label)
	}

	return result, nil
}

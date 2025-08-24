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
	builder := squirrel.Select(model.BookLabelColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("book_labels").
		Where(squirrel.Eq{
			"book_id": bookID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	result := make([]core.BookLabel, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		label := core.BookLabel{}

		err := rows.Scan(model.BookLabelScanner(&label))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, label)
	}

	return result, nil
}

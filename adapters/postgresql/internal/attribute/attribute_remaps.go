package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error) {
	builder := squirrel.Select(model.AttributeRemapColumns()...).
		From("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.AttributeRemap, 0)

	for rows.Next() {
		ar := core.AttributeRemap{}

		err = rows.Scan(model.AttributeRemapScanner(&ar))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, ar)
	}

	return result, nil
}

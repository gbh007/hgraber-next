package label

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error) {
	builder := squirrel.Select(
		"name",
		"description",
		"values",
		"created_at",
		"updated_at",
	).
		From("label_presets").
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

	result := make([]core.BookLabelPreset, 0, 10) //nolint:mnd // оптимизация

	for rows.Next() {
		var (
			preset      core.BookLabelPreset
			updatedAt   sql.NullTime
			description sql.NullString
		)

		err = rows.Scan(
			&preset.Name,
			&description,
			&preset.Values,
			&preset.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		preset.Description = description.String
		preset.UpdatedAt = updatedAt.Time

		result = append(result, preset)
	}

	return result, nil
}

package label

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error) {
	builder := squirrel.Select(
		"name",
		"description",
		"values",
		"created_at",
		"updated_at",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("label_presets").
		Where(squirrel.Eq{
			"name": name,
		}).
		Limit(1)

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	var (
		preset      core.BookLabelPreset
		updatedAt   sql.NullTime
		description sql.NullString
	)

	err := row.Scan(
		&preset.Name,
		&description,
		&preset.Values,
		&preset.CreatedAt,
		&updatedAt,
	)
	if err != nil { // TODO: err no rows
		return core.BookLabelPreset{}, fmt.Errorf("scan row: %w", err)
	}

	preset.Description = description.String
	preset.UpdatedAt = updatedAt.Time

	return preset, nil
}

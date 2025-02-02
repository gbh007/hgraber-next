package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/internal/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/internal/entities"
)

func (d *Database) InsertLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error {
	builder := squirrel.Insert("label_presets").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"name":        preset.Name,
			"description": model.StringToDB(preset.Description),
			"values":      preset.Values,
			"created_at":  preset.CreatedAt,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) UpdateLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error {
	builder := squirrel.Update("label_presets").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"description": model.StringToDB(preset.Description),
			"values":      preset.Values,
			"updated_at":  model.TimeToDB(preset.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"name": preset.Name,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeleteLabelPreset(ctx context.Context, name string) error {
	builder := squirrel.Delete("label_presets").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"name": name,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) LabelPresets(ctx context.Context) ([]entities.BookLabelPreset, error) {
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

	d.squirrelDebugLog(ctx, query, args)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]entities.BookLabelPreset, 0, 10)

	for rows.Next() {
		var (
			preset      entities.BookLabelPreset
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

func (d *Database) LabelPreset(ctx context.Context, name string) (entities.BookLabelPreset, error) {
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

	query, args, err := builder.ToSql()
	if err != nil {
		return entities.BookLabelPreset{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	row := d.pool.QueryRow(ctx, query, args...)

	var (
		preset      entities.BookLabelPreset
		updatedAt   sql.NullTime
		description sql.NullString
	)

	err = row.Scan(
		&preset.Name,
		&description,
		&preset.Values,
		&preset.CreatedAt,
		&updatedAt,
	)
	if err != nil { // TODO: err no rows
		return entities.BookLabelPreset{}, fmt.Errorf("scan row: %w", err)
	}

	preset.Description = description.String
	preset.UpdatedAt = updatedAt.Time

	return preset, nil
}

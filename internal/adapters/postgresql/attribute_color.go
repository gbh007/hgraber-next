package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (d *Database) InsertAttributeColor(ctx context.Context, color entities.AttributeColor) error {
	builder := squirrel.Insert("attribute_colors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"attr":             color.Code,
			"value":            color.Value,
			"text_color":       color.TextColor,
			"background_color": color.BackgroundColor,
			"created_at":       color.CreatedAt,
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

func (d *Database) UpdateAttributeColor(ctx context.Context, color entities.AttributeColor) error {
	builder := squirrel.Update("attribute_colors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"text_color":       color.TextColor,
			"background_color": color.BackgroundColor,
		}).
		Where(squirrel.Eq{
			"attr":  color.Code,
			"value": color.Value,
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

func (d *Database) DeleteAttributeColor(ctx context.Context, code, value string) error {
	builder := squirrel.Delete("attribute_colors").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
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

func (d *Database) AttributeColors(ctx context.Context) ([]entities.AttributeColor, error) {
	builder := squirrel.Select(
		"attr",
		"value",
		"text_color",
		"background_color",
		"created_at",
	).
		From("attribute_colors").
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

	result := make([]entities.AttributeColor, 0, 10)

	for rows.Next() {
		color := entities.AttributeColor{}

		err = rows.Scan(
			&color.Code,
			&color.Value,
			&color.TextColor,
			&color.BackgroundColor,
			&color.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, color)
	}

	return result, nil
}

func (d *Database) AttributeColor(ctx context.Context, code, value string) (entities.AttributeColor, error) {
	builder := squirrel.Select(
		"attr",
		"value",
		"text_color",
		"background_color",
		"created_at",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("attribute_colors").
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return entities.AttributeColor{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	row := d.pool.QueryRow(ctx, query, args...)

	color := entities.AttributeColor{}

	err = row.Scan(
		&color.Code,
		&color.Value,
		&color.TextColor,
		&color.BackgroundColor,
		&color.CreatedAt,
	)
	if err != nil { // TODO: err no rows
		return entities.AttributeColor{}, fmt.Errorf("scan row: %w", err)
	}

	return color, nil
}

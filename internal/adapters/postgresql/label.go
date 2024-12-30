package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (d *Database) SetLabel(ctx context.Context, label entities.BookLabel) error {
	builder := squirrel.Insert("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"book_id":     label.BookID.String(),
			"page_number": label.PageNumber,
			"name":        label.Name,
			"value":       label.Value,
			"create_at":   label.CreateAt,
		}).
		Suffix(`ON CONFLICT (book_id, page_number, name) DO UPDATE SET value = EXCLUDED.value`)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeleteLabel(ctx context.Context, label entities.BookLabel) error {
	builder := squirrel.Delete("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id":     label.BookID.String(),
			"page_number": label.PageNumber,
			"name":        label.Name,
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

func (d *Database) Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error) {
	raw := make([]model.BookLabel, 0)

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From("book_labels").
		Where(squirrel.Eq{
			"book_id": bookID.String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	err = d.db.SelectContext(ctx, &raw, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	result, err := pkg.MapWithError(raw, func(a model.BookLabel) (entities.BookLabel, error) {
		return a.ToEntity()
	})
	if err != nil {
		return nil, fmt.Errorf("convert: %w", err)
	}

	return result, nil
}

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

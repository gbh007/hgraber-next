package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

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

func (d *Database) ReplaceLabels(ctx context.Context, bookID uuid.UUID, labels []entities.BookLabel) error {
	builder := squirrel.Insert("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"page_number",
			"name",
			"value",
			"create_at",
		).
		Suffix(`ON CONFLICT (book_id, page_number, name) DO UPDATE SET value = EXCLUDED.value`)

	for _, label := range labels {
		builder = builder.Values(
			label.BookID.String(),
			label.PageNumber,
			label.Name,
			label.Value,
			label.CreateAt,
		)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			d.logger.ErrorContext(
				ctx, "rollback ReplaceLabels tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM book_labels WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete old labels: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

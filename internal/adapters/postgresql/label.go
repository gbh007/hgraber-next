package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (d *Database) SetLabel(ctx context.Context, label entities.BookLabel) error {
	builder := squirrel.Insert("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"page_number",
			"name",
			"value",
			"create_at",
		).
		Values(
			label.BookID.String(),
			model.Int32ToDB(label.PageNumber),
			label.Name,
			label.Value,
			label.CreateAt,
		).Suffix(`ON CONFLICT (book_id, page_number, name) DO UPDATE SET value = EXCLUDED.value`)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)

	_, err = d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error) {
	raw := make([]model.BookLabel, 0)

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From("book_labels").Where(squirrel.Eq{
		"book_id": bookID.String(),
	})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)

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

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

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (d *Database) SetLabel(ctx context.Context, label core.BookLabel) error {
	builder := squirrel.Insert("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"book_id":     label.BookID,
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

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeleteLabel(ctx context.Context, label core.BookLabel) error {
	builder := squirrel.Delete("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id":     label.BookID,
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

func (d *Database) DeleteBookLabels(ctx context.Context, bookID uuid.UUID) error {
	builder := squirrel.Delete("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id": bookID,
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

func (d *Database) Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error) {
	raw := make([]model.BookLabel, 0)

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From("book_labels").
		Where(squirrel.Eq{
			"book_id": bookID,
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

	result, err := pkg.MapWithError(raw, func(a model.BookLabel) (core.BookLabel, error) {
		return a.ToEntity()
	})
	if err != nil {
		return nil, fmt.Errorf("convert: %w", err)
	}

	return result, nil
}

func (d *Database) ReplaceLabels(ctx context.Context, bookID uuid.UUID, labels []core.BookLabel) error {
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
			bookID,
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

func (d *Database) SetLabels(ctx context.Context, labels []core.BookLabel) error {
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
			label.BookID,
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

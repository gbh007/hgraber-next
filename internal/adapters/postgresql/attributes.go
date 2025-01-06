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
)

// TODO: объединить с потребителем
func (d *Database) bookAttributes(ctx context.Context, bookID uuid.UUID) ([]*model.BookAttribute, error) {
	raw := make([]*model.BookAttribute, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM book_attributes WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (d *Database) BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	attributes, err := d.bookAttributes(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("get attributes: %w", err)
	}

	out := make(map[string][]string, entities.PossibleAttributeCount)

	for _, attribute := range attributes {
		out[attribute.Attr] = append(out[attribute.Attr], attribute.Value)
	}

	return out, nil
}

func (d *Database) UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.ExecContext(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("delete old attributes: %w", err)
	}

	builder := squirrel.Insert("book_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"attr",
			"value",
		)

	for code, values := range attributes {
		for _, value := range values {
			builder = builder.Values(bookID.String(), code, value)
		}
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *Database) DeleteBookAttributes(ctx context.Context, bookID uuid.UUID) error {
	_, err := d.pool.Exec(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) AttributesCount(ctx context.Context) ([]entities.AttributeVariant, error) {
	rows, err := d.pool.Query(ctx, `SELECT COUNT(*), attr, value FROM book_attributes GROUP BY attr, value;`)
	if err != nil {
		return nil, fmt.Errorf("get attributes count: %w", err)
	}

	defer rows.Close()

	result := make([]entities.AttributeVariant, 0, 100) // Берем изначальный запас емкости побольше

	for rows.Next() {
		var (
			count int
			code  string
			value string
		)

		err := rows.Scan(&count, &code, &value)
		if err != nil {
			return nil, fmt.Errorf("get attributes count: scan row: %w", err)
		}

		result = append(result, entities.AttributeVariant{
			Code:  code,
			Value: value,
			Count: count,
		})
	}

	return result, nil
}

func (d *Database) Attributes(ctx context.Context) ([]entities.Attribute, error) {
	builder := squirrel.Select(
		"code",
		"name",
		"plural_name",
		"\"order\"",
		"description",
	).
		From("attributes").
		PlaceholderFormat(squirrel.Dollar).
		OrderBy("\"order\"")

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

	result := make([]entities.Attribute, 0, entities.PossibleAttributeCount)

	for rows.Next() {
		attribute := entities.Attribute{}
		description := sql.NullString{}

		err = rows.Scan(
			&attribute.Code,
			&attribute.Name,
			&attribute.PluralName,
			&attribute.Order,
			&description,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		attribute.Description = description.String

		result = append(result, attribute)
	}

	return result, nil
}

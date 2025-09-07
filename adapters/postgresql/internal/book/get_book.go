package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error) {
	builder := squirrel.Select(model.BookColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.Eq{
			"id": bookID,
		}).
		Limit(1)

	query, args := builder.MustSql()

	book := core.Book{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(model.BookScanner(&book))

	if errors.Is(err, sql.ErrNoRows) {
		return core.Book{}, core.BookNotFoundError
	}

	if err != nil {
		return core.Book{}, fmt.Errorf("exec query: %w", err)
	}

	return book, nil
}

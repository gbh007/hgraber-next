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
	bookTable := model.BookTable

	builder := squirrel.Select(bookTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(bookTable.Name()).
		Where(squirrel.Eq{
			bookTable.ColumnID(): bookID,
		}).
		Limit(1)

	query, args := builder.MustSql()

	book := core.Book{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(bookTable.Scanner(&book))

	if errors.Is(err, sql.ErrNoRows) {
		return core.Book{}, core.ErrBookNotFound
	}

	if err != nil {
		return core.Book{}, fmt.Errorf("exec query: %w", err)
	}

	return book, nil
}

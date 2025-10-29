package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

// TODO: добавить лимиты
func (repo *BookRepo) UnprocessedBooks(ctx context.Context) ([]core.Book, error) {
	bookTable := model.BookTable

	builder := squirrel.Select(bookTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(bookTable.Name()).
		Where(squirrel.And{
			squirrel.Or{
				squirrel.Expr(bookTable.ColumnName() + " IS NULL"),
				squirrel.Expr(bookTable.ColumnPageCount() + " IS NULL"),
				squirrel.Eq{
					bookTable.ColumnAttributesParsed(): false,
				},
			},
			squirrel.Expr(bookTable.ColumnOriginURL() + " IS NOT NULL"),
			squirrel.Eq{
				bookTable.ColumnDeleted():   false,
				bookTable.ColumnIsRebuild(): false,
			},
		})

	query, args := builder.MustSql()

	result := make([]core.Book, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		book := core.Book{}

		err := rows.Scan(bookTable.Scanner(&book))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, book)
	}

	return result, nil
}

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
	builder := squirrel.Select(model.BookColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.And{
			squirrel.Or{
				squirrel.Expr(`name IS NULL`),
				squirrel.Expr(`page_count IS NULL`),
				squirrel.Eq{
					"attributes_parsed": false,
				},
			},
			squirrel.Expr(`origin_url IS NOT NULL`),
			squirrel.Eq{
				"deleted":    false,
				"is_rebuild": false,
			},
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	result := make([]core.Book, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		book := core.Book{}

		err := rows.Scan(model.BookScanner(&book))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, book)
	}

	return result, nil
}

package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) UpdateBookDeletion(ctx context.Context, book core.Book) error {
	bookTable := model.BookTable

	builder := squirrel.Update(bookTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			bookTable.ColumnDeleted():   book.Deleted,
			bookTable.ColumnDeletedAt(): model.TimeToDB(book.DeletedAt),
		}).
		Where(squirrel.Eq{
			bookTable.ColumnID(): book.ID,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("storage: exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrBookNotFound
	}

	return nil
}

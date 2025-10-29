package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) UpdateBook(ctx context.Context, book core.Book) error {
	bookTable := model.BookTable

	builder := squirrel.Update(bookTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			bookTable.ColumnName():             model.StringToDB(book.Name),
			bookTable.ColumnOriginURL():        model.URLToDB(book.OriginURL),
			bookTable.ColumnPageCount():        model.Int32ToDB(book.PageCount),
			bookTable.ColumnAttributesParsed(): book.AttributesParsed,
			bookTable.ColumnVerified():         book.Verified,
			bookTable.ColumnVerifiedAt():       model.TimeToDB(book.VerifiedAt),
			bookTable.ColumnIsRebuild():        book.IsRebuild,
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

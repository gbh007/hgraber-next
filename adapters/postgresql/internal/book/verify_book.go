package book

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error {
	bookTable := model.BookTable

	builder := squirrel.Update(bookTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			bookTable.ColumnVerified():   verified,
			bookTable.ColumnVerifiedAt(): model.TimeToDB(verifiedAt),
		}).
		Where(squirrel.Eq{
			bookTable.ColumnID(): bookID,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrBookNotFound
	}

	return nil
}

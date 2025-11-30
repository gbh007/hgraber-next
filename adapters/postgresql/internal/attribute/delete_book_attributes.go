package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *AttributeRepo) DeleteBookAttributes(ctx context.Context, bookID uuid.UUID) error {
	bookAttributeTable := model.BookAttributeTable

	query, args := squirrel.
		Delete(bookAttributeTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			bookAttributeTable.ColumnBookID(): bookID,
		}).
		MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

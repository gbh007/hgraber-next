package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	bookAttributeTable := model.BookAttributeTable

	query, args := squirrel.Select(
		bookAttributeTable.ColumnAttr(),
		bookAttributeTable.ColumnValue(),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(bookAttributeTable.Name()).
		Where(squirrel.Eq{
			bookAttributeTable.ColumnBookID(): bookID,
		}).
		MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select rows: %w", err)
	}

	defer rows.Close()

	out := make(map[string][]string, core.PossibleAttributeCount)

	for rows.Next() {
		var (
			code  string
			value string
		)

		err = rows.Scan(&code, &value)
		if err != nil {
			return nil, fmt.Errorf("scan rows: %w", err)
		}

		out[code] = append(out[code], value)
	}

	return out, nil
}

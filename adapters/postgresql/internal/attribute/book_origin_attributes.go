package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	bookOriginAttributeTable := model.BookOriginAttributeTable

	query, args := squirrel.Select(
		bookOriginAttributeTable.ColumnAttr(),
		bookOriginAttributeTable.ColumnValues(),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(bookOriginAttributeTable.Name()).
		Where(squirrel.Eq{
			bookOriginAttributeTable.ColumnBookID(): bookID,
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
			code   string
			values []string
		)

		err = rows.Scan(&code, &values)
		if err != nil {
			return nil, fmt.Errorf("scan rows: %w", err)
		}

		out[code] = values
	}

	return out, nil
}

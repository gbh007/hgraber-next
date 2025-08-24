package attribute

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	rows, err := repo.Pool.Query(ctx, `SELECT attr, value FROM book_attributes WHERE book_id = $1;`, bookID)
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

package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) InsertAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	builder := squirrel.Insert("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"attr":       ar.Code,
			"value":      ar.Value,
			"to_attr":    model.StringToDB(ar.ToCode),
			"to_value":   model.StringToDB(ar.ToValue),
			"created_at": ar.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

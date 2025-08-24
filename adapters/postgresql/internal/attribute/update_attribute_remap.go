package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) UpdateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	builder := squirrel.Update("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"to_attr":    model.StringToDB(ar.ToCode),
			"to_value":   model.StringToDB(ar.ToValue),
			"updated_at": model.TimeToDB(ar.UpdateAt),
		}).
		Where(squirrel.Eq{
			"attr":  ar.Code,
			"value": ar.Value,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

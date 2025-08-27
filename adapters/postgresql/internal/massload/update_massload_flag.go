package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error {
	builder := squirrel.Update("massload_flags").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":             flag.Name,
			"description":      model.StringToDB(flag.Description),
			"order_weight":     flag.OrderWeight,
			"text_color":       model.StringToDB(flag.TextColor),
			"background_color": model.StringToDB(flag.BackgroundColor),
		}).
		Where(squirrel.Eq{
			"code": flag.Code,
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

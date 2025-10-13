package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error {
	table := model.MassloadFlagTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnName():            flag.Name,
			table.ColumnDescription():     model.StringToDB(flag.Description),
			table.ColumnOrderWeight():     flag.OrderWeight,
			table.ColumnTextColor():       model.StringToDB(flag.TextColor),
			table.ColumnBackgroundColor(): model.StringToDB(flag.BackgroundColor),
		}).
		Where(squirrel.Eq{
			table.ColumnCode(): flag.Code,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

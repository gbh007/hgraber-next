package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) CreateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error {
	table := model.MassloadFlagTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnCode():            flag.Code,
			table.ColumnName():            flag.Name,
			table.ColumnDescription():     model.StringToDB(flag.Description),
			table.ColumnOrderWeight():     flag.OrderWeight,
			table.ColumnTextColor():       model.StringToDB(flag.TextColor),
			table.ColumnBackgroundColor(): model.StringToDB(flag.BackgroundColor),
			table.ColumnCreatedAt():       flag.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

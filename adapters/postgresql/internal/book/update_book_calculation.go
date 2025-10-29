package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) UpdateBookCalculation(ctx context.Context, id uuid.UUID, calc core.BookCalc) error {
	bookTable := model.BookTable

	builder := squirrel.Update(bookTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			bookTable.ColumnCalcPageCount():     model.NilInt64ToDB(calc.CalcPageCount),
			bookTable.ColumnCalcFileCount():     model.NilInt64ToDB(calc.CalcFileCount),
			bookTable.ColumnCalcDeadHashCount(): model.NilInt64ToDB(calc.CalcDeadHashCount),
			bookTable.ColumnCalcPageSize():      model.NilInt64ToDB(calc.CalcPageSize),
			bookTable.ColumnCalcFileSize():      model.NilInt64ToDB(calc.CalcFileSize),
			bookTable.ColumnCalcDeadHashSize():  model.NilInt64ToDB(calc.CalcDeadHashSize),
			bookTable.ColumnCalculatedAt():      model.TimeToDB(calc.CalculatedAt),
			bookTable.ColumnCalcAvgPageSize():   model.NilInt64ToDB(calc.CalcAvgPageSize),
		}).
		Where(squirrel.Eq{
			bookTable.ColumnID(): id,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("storage: exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrBookNotFound
	}

	return nil
}

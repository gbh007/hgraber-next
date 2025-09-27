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
	builder := squirrel.Update("books").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"calc_page_count":      model.NilInt64ToDB(calc.CalcPageCount),
			"calc_file_count":      model.NilInt64ToDB(calc.CalcFileCount),
			"calc_dead_hash_count": model.NilInt64ToDB(calc.CalcDeadHashCount),
			"calc_page_size":       model.NilInt64ToDB(calc.CalcPageSize),
			"calc_file_size":       model.NilInt64ToDB(calc.CalcFileSize),
			"calc_dead_hash_size":  model.NilInt64ToDB(calc.CalcDeadHashSize),
			"calculated_at":        model.TimeToDB(calc.CalculatedAt),
		}).
		Where(squirrel.Eq{
			"id": id,
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

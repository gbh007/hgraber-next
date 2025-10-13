package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadExternalLinkCounts(
	ctx context.Context,
	id int,
	link massloadmodel.ExternalLink,
) error {
	table := model.MassloadExternalLinkTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnBooksAhead():    model.NilInt64ToDB(link.BooksAhead),
			table.ColumnNewBooks():      model.NilInt64ToDB(link.NewBooks),
			table.ColumnExistingBooks(): model.NilInt64ToDB(link.ExistingBooks),
			table.ColumnUpdatedAt():     link.UpdatedAt,
		}).
		Where(squirrel.Eq{
			table.ColumnMassloadID(): id,
			table.ColumnURL():        link.URL.String(),
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

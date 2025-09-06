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
	builder := squirrel.Update("massload_external_links").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"books_ahead":    model.NilInt64ToDB(link.BooksAhead),
			"new_books":      model.NilInt64ToDB(link.NewBooks),
			"existing_books": model.NilInt64ToDB(link.ExistingBooks),
			"updated_at":     link.UpdatedAt,
		}).
		Where(squirrel.Eq{
			"massload_id": id,
			"url":         link.URL.String(),
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

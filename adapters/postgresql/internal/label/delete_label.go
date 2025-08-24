package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) DeleteLabel(ctx context.Context, label core.BookLabel) error {
	builder := squirrel.Delete("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id":     label.BookID,
			"page_number": label.PageNumber,
			"name":        label.Name,
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

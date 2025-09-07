package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) SetLabel(ctx context.Context, label core.BookLabel) error {
	builder := squirrel.Insert("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"book_id":     label.BookID,
			"page_number": label.PageNumber,
			"name":        label.Name,
			"value":       label.Value,
			"create_at":   label.CreateAt,
		}).
		Suffix(`ON CONFLICT (book_id, page_number, name) DO UPDATE SET value = EXCLUDED.value`)

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

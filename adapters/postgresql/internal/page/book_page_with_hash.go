package page

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPageWithHash(
	ctx context.Context,
	bookID uuid.UUID,
	pageNumber int,
) (core.PageWithHash, error) {
	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"p.book_id":     bookID,
			"p.page_number": pageNumber,
		}).
		Limit(1)

	query, args := builder.MustSql()

	page := core.PageWithHash{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(model.PageWithHashScanner(&page))

	if errors.Is(err, sql.ErrNoRows) {
		return core.PageWithHash{}, core.ErrPageNotFound
	}

	if err != nil {
		return core.PageWithHash{}, fmt.Errorf("exec query: %w", err)
	}

	return page, nil
}

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
	pageTable := model.PageTable.WithPrefix("p")
	fileTable := model.FileTable.WithPrefix("f")
	pageWithHashTable := model.NewPageWithHash(pageTable, fileTable)

	builder := squirrel.Select(pageWithHashTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(pageTable.NameAlter()).
		LeftJoin(pageWithHashTable.JoinString()).
		Where(squirrel.Eq{
			pageTable.ColumnBookID():     bookID,
			pageTable.ColumnPageNumber(): pageNumber,
		}).
		Limit(1)

	query, args := builder.MustSql()

	page := core.PageWithHash{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(pageWithHashTable.Scanner(&page))

	if errors.Is(err, sql.ErrNoRows) {
		return core.PageWithHash{}, core.ErrPageNotFound
	}

	if err != nil {
		return core.PageWithHash{}, fmt.Errorf("exec query: %w", err)
	}

	return page, nil
}

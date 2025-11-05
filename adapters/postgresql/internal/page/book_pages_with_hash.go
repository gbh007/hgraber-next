package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error) {
	pageTable := model.PageTable.WithPrefix("p")
	fileTable := model.FileTable.WithPrefix("f")
	pageWithHashTable := model.NewPageWithHash(pageTable, fileTable)

	builder := squirrel.Select(pageWithHashTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(pageTable.NameAlter()).
		LeftJoin(pageWithHashTable.JoinString()).
		Where(squirrel.Eq{
			pageTable.ColumnBookID(): bookID,
		}).
		OrderBy(pageTable.ColumnPageNumber())

	query, args := builder.MustSql()

	out := make([]core.PageWithHash, 0, core.AvgPageCountInBook)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.PageWithHash{}

		err := rows.Scan(pageWithHashTable.Scanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}

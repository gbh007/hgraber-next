package page

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) PageSizeByAuthor(ctx context.Context) (map[string]core.SizeWithCount, error) {
	pageTable := model.PageTable.WithPrefix("p")
	fileTable := model.FileTable.WithPrefix("f")
	bookAttributeTable := model.BookAttributeTable.WithPrefix("a")

	builder := squirrel.Select("COUNT(*)", bookAttributeTable.ColumnValue(), "SUM("+fileTable.ColumnSize()+")").
		PlaceholderFormat(squirrel.Dollar).
		From(bookAttributeTable.NameAlter()).
		InnerJoin(model.JoinBookAttributePage(bookAttributeTable, pageTable)).
		InnerJoin(model.JoinPageAndFile(pageTable, fileTable)).
		Where(squirrel.Eq{
			bookAttributeTable.ColumnAttr(): "author",
		}).
		GroupBy(bookAttributeTable.ColumnValue())

	query, args := builder.MustSql()

	out := make(map[string]core.SizeWithCount, 100) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count sql.NullInt64
			size  sql.NullInt64
			name  string
		)

		err = rows.Scan(&count, &name, &size)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[name] = core.SizeWithCount{
			Count: count.Int64,
			Size:  size.Int64,
		}
	}

	return out, nil
}

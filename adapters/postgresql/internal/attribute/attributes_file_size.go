package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributesFileSize(
	ctx context.Context,
	attrs map[string][]string,
) (core.SizeWithCount, error) {
	fileTable := model.FileTable.WithPrefix("f")
	pageTable := model.PageTable.WithPrefix("p")
	bookAttributeTable := model.BookAttributeTable.WithPrefix("a")

	whereCond := squirrel.Or{}

	for code, values := range attrs {
		if len(values) == 0 {
			continue
		}

		whereCond = append(whereCond, squirrel.Eq{
			bookAttributeTable.ColumnAttr():  code,
			bookAttributeTable.ColumnValue(): values,
		})
	}

	if len(whereCond) == 0 {
		return core.SizeWithCount{}, errors.New("incorrect condition: empty attributes")
	}

	subBuilder := squirrel.Select(
		fileTable.ColumnSize()+" AS \"size\"",
		fileTable.ColumnMd5Sum(),
		fileTable.ColumnSha256Sum(),
	).
		// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
		PlaceholderFormat(squirrel.Question).
		From(fileTable.NameAlter()).
		InnerJoin(model.JoinFileAndPage(fileTable, pageTable)).
		InnerJoin(model.JoinPageAndBookAttribute(pageTable, bookAttributeTable)).
		Where(whereCond).
		GroupBy(
			fileTable.ColumnSize(),
			fileTable.ColumnMd5Sum(),
			fileTable.ColumnSha256Sum(),
		)

	builder := squirrel.Select(`SUM(uf."size")`, `COUNT(*)`).
		PlaceholderFormat(squirrel.Dollar).
		FromSelect(subBuilder, "uf")

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	var size, count sql.NullInt64

	err := row.Scan(&size, &count)
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("exec query: %w", err)
	}

	return core.SizeWithCount{
		Count: count.Int64,
		Size:  size.Int64,
	}, nil
}

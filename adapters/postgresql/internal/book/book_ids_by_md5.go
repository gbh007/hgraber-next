package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *BookRepo) BookIDsByMD5(ctx context.Context, md5Sums []string) ([]uuid.UUID, error) {
	bookTable := model.BookTable.WithPrefix("b")
	fileTable := model.FileTable.WithPrefix("f")
	pageTable := model.PageTable.WithPrefix("p")

	builder := squirrel.Select(bookTable.ColumnID()).
		PlaceholderFormat(squirrel.Dollar).
		From(bookTable.NameAlter()).
		InnerJoin(model.JoinBookAndPage(bookTable, pageTable)).
		InnerJoin(model.JoinPageAndFile(pageTable, fileTable)).
		Where(squirrel.Eq{
			fileTable.ColumnMd5Sum(): md5Sums,
		}).
		GroupBy(bookTable.ColumnID())

	query, args := builder.MustSql()

	result := []uuid.UUID{}

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		id := uuid.UUID{}

		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, id)
	}

	return result, nil
}

package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *BookRepo) BookIDsByMD5(ctx context.Context, md5Sums []string) ([]uuid.UUID, error) {
	bookTable := model.BookTable
	fileTable := model.FileTable
	pageTable := model.PageTable

	builder := squirrel.Select("b." + bookTable.ColumnID()).
		PlaceholderFormat(squirrel.Dollar).
		From(bookTable.Name() + " b").
		InnerJoin(pageTable.Name() + " p ON p." + pageTable.ColumnBookID() + " = b." + bookTable.ColumnID()).
		InnerJoin(fileTable.Name() + " f ON f." + fileTable.ColumnID() + " = p." + pageTable.ColumnFileID()).
		Where(squirrel.Eq{
			"f." + fileTable.ColumnMd5Sum(): md5Sums,
		}).
		GroupBy("b." + bookTable.ColumnID())

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

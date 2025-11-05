package page

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPagesCountByHash(ctx context.Context, hash core.FileHash) (int64, error) {
	pageTable := model.PageTable.WithPrefix("p")
	fileTable := model.FileTable.WithPrefix("f")

	builder := squirrel.Select("COUNT(*)").
		PlaceholderFormat(squirrel.Dollar).
		From(pageTable.NameAlter()).
		LeftJoin(model.JoinPageAndFile(pageTable, fileTable)).
		Where(squirrel.Eq{
			fileTable.ColumnMd5Sum():    hash.Md5Sum,
			fileTable.ColumnSha256Sum(): hash.Sha256Sum,
			fileTable.ColumnSize():      hash.Size,
		})

	query, args := builder.MustSql()

	count := sql.NullInt64{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get count: %w", err)
	}

	return count.Int64, nil
}

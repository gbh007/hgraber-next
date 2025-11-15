package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) RemoveDeletedPagesByHash(ctx context.Context, hash core.FileHash) error {
	deletedPageTable := model.DeletedPageTable

	builder := squirrel.Delete(deletedPageTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			deletedPageTable.ColumnMd5Sum():    hash.Md5Sum,
			deletedPageTable.ColumnSha256Sum(): hash.Sha256Sum,
			deletedPageTable.ColumnSize():      hash.Size,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

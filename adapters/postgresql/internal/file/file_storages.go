package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (repo *FileRepo) FileStorages(ctx context.Context) ([]fsmodel.FileStorageSystem, error) {
	builder := squirrel.Select(model.FileStorageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("file_storages")

	query, args := builder.MustSql()

	out := make([]fsmodel.FileStorageSystem, 0, 10) //nolint:mnd // будет исправлено позднее

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		fs := fsmodel.FileStorageSystem{}

		err := rows.Scan(model.FileStorageScanner(&fs))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, fs)
	}

	return out, nil
}

package file

import (
	"context"
	"fmt"
	"strings"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) DuplicatedFiles(ctx context.Context) ([]core.File, error) {
	// TODO: переделать на squirrel
	query := `SELECT ` + strings.Join(model.StringsPrefix(model.FileColumns(), "f."), ", ") + `
FROM (
        SELECT COUNT(*) AS c, md5_sum, sha256_sum
        FROM files
        GROUP BY
            md5_sum, sha256_sum
        HAVING
            COUNT(*) > 1
    ) AS t
    INNER join files AS f ON f.md5_sum = t.md5_sum
    AND f.sha256_sum = t.sha256_sum ORDER BY f.id;`

	result := make([]core.File, 0)

	rows, err := repo.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		file := core.File{}

		err := rows.Scan(model.FileScanner(&file))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, file)
	}

	return result, nil
}

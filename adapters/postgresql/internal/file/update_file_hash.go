package file

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) UpdateFileHash(ctx context.Context, id uuid.UUID, md5Sum, sha256Sum string, size int64) error {
	res, err := repo.Pool.Exec(
		ctx,
		`UPDATE files SET md5_sum = $2, sha256_sum = $3, "size" = $4 WHERE id = $1`,
		id, model.StringToDB(md5Sum), model.StringToDB(sha256Sum), model.Int64ToDB(size),
	)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.FileNotFoundError
	}

	return nil
}

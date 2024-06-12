package postgresql

import (
	"context"
	"database/sql"
	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"

	"github.com/google/uuid"
)

func (d *Database) GetUnHashedFiles(ctx context.Context) []entities.File {
	raw := make([]*model.File, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM files WHERE md5_sum IS NULL OR sha256_sum IS NULL OR "size" IS NULL;`)
	if err != nil {
		d.logger.ErrorContext(ctx, err.Error())

		return []entities.File{}
	}

	out := make([]entities.File, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			d.logger.ErrorContext(ctx, err.Error())

			return []entities.File{}
		}
	}

	return out
}

func (d *Database) UpdateFileHash(ctx context.Context, id uuid.UUID, md5Sum, sha256Sum string, size int64) error {
	res, err := d.db.ExecContext(
		ctx,
		`UPDATE files SET md5_sum = $2, sha256_sum = $3, "size" = $4 WHERE id = $1`,
		id.String(), model.StringToDB(md5Sum), model.StringToDB(sha256Sum), sql.NullInt64{Int64: size, Valid: size > 0},
	)
	if err != nil {
		return err
	}

	if !d.isApply(ctx, res) {
		return entities.FileNotFoundError
	}

	return nil
}

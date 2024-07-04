package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
)

func (d *Database) GetUnHashedFiles(ctx context.Context) ([]entities.File, error) {
	raw := make([]*model.File, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM files WHERE md5_sum IS NULL OR sha256_sum IS NULL OR "size" IS NULL;`)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	out := make([]entities.File, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert %s: %w", v.ID, err)
		}
	}

	return out, nil
}

func (d *Database) DuplicatedFiles(ctx context.Context) ([]entities.File, error) {
	raw := make([]*model.File, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT f.*
FROM (
        SELECT COUNT(*) AS c, md5_sum, sha256_sum
        FROM files
        GROUP BY
            md5_sum, sha256_sum
        HAVING
            COUNT(*) > 1
    ) AS t
    INNER join files AS f ON f.md5_sum = t.md5_sum
    AND f.sha256_sum = t.sha256_sum ORDER BY f.id;`)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	out := make([]entities.File, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert %s: %w", v.ID, err)
		}
	}

	return out, nil
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

func (d *Database) NewFile(ctx context.Context, file entities.File) error {
	builder := squirrel.Insert("files").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"id",
			"filename",
			"ext",
			"create_at",
		).
		Values(
			file.ID.String(),
			file.Filename,
			file.Ext,
			file.CreateAt,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)

	_, err = d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) ReplaceFile(ctx context.Context, oldFileID, newFileID uuid.UUID) error {
	builder := squirrel.Update("pages").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(
			map[string]interface{}{
				"file_id": newFileID.String(),
			},
		).
		Where(squirrel.Eq{
			"file_id": oldFileID.String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DetachedFiles(ctx context.Context) ([]entities.File, error) {
	raw := make([]*model.File, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT *
FROM files
WHERE
    id NOT IN (
        SELECT file_id
        FROM pages
    );`)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	out := make([]entities.File, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert %s: %w", v.ID, err)
		}
	}

	return out, nil
}

func (d *Database) DeleteFile(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("files").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id.String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if !d.isApply(ctx, res) {
		return entities.FileNotFoundError
	}

	return nil
}

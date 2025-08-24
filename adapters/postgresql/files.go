package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (d *Database) DuplicatedFiles(ctx context.Context) ([]core.File, error) {
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

	rows, err := d.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
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

func (d *Database) UpdateFileHash(ctx context.Context, id uuid.UUID, md5Sum, sha256Sum string, size int64) error {
	res, err := d.pool.Exec(
		ctx,
		`UPDATE files SET md5_sum = $2, sha256_sum = $3, "size" = $4 WHERE id = $1`,
		id, model.StringToDB(md5Sum), model.StringToDB(sha256Sum), sql.NullInt64{Int64: size, Valid: size > 0},
	)
	if err != nil {
		return err
	}

	if res.RowsAffected() < 1 {
		return core.FileNotFoundError
	}

	return nil
}

func (d *Database) NewFile(ctx context.Context, file core.File) error {
	builder := squirrel.Insert("files").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"id":        file.ID,
			"filename":  file.Filename,
			"ext":       file.Ext,
			"create_at": file.CreateAt,
			"fs_id":     file.FSID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) ReplaceFile(ctx context.Context, oldFileID, newFileID uuid.UUID) error {
	builder := squirrel.Update("pages").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"file_id": newFileID,
		}).
		Where(squirrel.Eq{
			"file_id": oldFileID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DetachedFiles(ctx context.Context) ([]core.File, error) {
	builder := squirrel.Select(model.FileColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Expr(`NOT EXISTS (SELECT 1 FROM pages WHERE pages.file_id = files.id)`))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]core.File, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
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

func (d *Database) FilesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.File, error) {
	builder := squirrel.Select(model.FileColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Eq{
			"md5_sum": md5Sums,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]core.File, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
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

func (d *Database) DeleteFile(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("files").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.FileNotFoundError
	}

	return nil
}

func (d *Database) UpdateFileInvalidData(ctx context.Context, fileID uuid.UUID, invalidData bool) error {
	builder := squirrel.Update("files").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"invalid_data": invalidData,
		}).
		Where(squirrel.Eq{
			"id": fileID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.FileNotFoundError
	}

	return nil
}

func (d *Database) UpdateFileFS(ctx context.Context, fileID, fsID uuid.UUID) error {
	builder := squirrel.Update("files").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"fs_id": fsID,
		}).
		Where(squirrel.Eq{
			"id": fileID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.FileNotFoundError
	}

	return nil
}

func (d *Database) File(ctx context.Context, id uuid.UUID) (core.File, error) {
	builder := squirrel.Select(model.FileColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.File{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	file := core.File{}

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.FileScanner(&file))
	if err != nil {
		return core.File{}, fmt.Errorf("exec: %w", err)
	}

	return file, nil
}

func (d *Database) FSFilesInfo(
	ctx context.Context,
	fsID uuid.UUID,
	onlyInvalidData, onlyDetached bool,
) (core.SizeWithCount, error) {
	builder := squirrel.Select(
		"COUNT(*)",
		"SUM(\"size\")",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Eq{
			"fs_id": fsID,
		})

	if onlyInvalidData {
		builder = builder.Where(squirrel.Eq{
			"invalid_data": true,
		})
	}

	if onlyDetached {
		builder = builder.Where(
			squirrel.Expr(`NOT EXISTS (SELECT 1 FROM pages WHERE file_id = files.id)`),
		)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	var count, size sql.NullInt64

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(&count, &size)
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("scan :%w", err)
	}

	return core.SizeWithCount{
		Count: count.Int64,
		Size:  size.Int64,
	}, nil
}

func (d *Database) FileIDsByFilter(ctx context.Context, filter fsmodel.FileFilter) ([]uuid.UUID, error) {
	builder := squirrel.Select("id").
		PlaceholderFormat(squirrel.Dollar).
		From("files")

	if filter.FSID != nil {
		builder = builder.Where(squirrel.Eq{
			"fs_id": *filter.FSID,
		})
	}

	if filter.BookID != nil || filter.PageNumber != nil {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From("pages").
			Where(squirrel.Expr(`file_id = files.id`))

		if filter.BookID != nil {
			subBuilder = subBuilder.Where(squirrel.Eq{
				"book_id": *filter.BookID,
			})
		}

		if filter.PageNumber != nil {
			subBuilder = subBuilder.Where(squirrel.Eq{
				"page_number": *filter.PageNumber,
			})
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build pages sub query: %w", err)
		}

		builder = builder.Where(squirrel.Expr(`EXISTS (`+subQuery+`)`, subArgs...))
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	defer rows.Close()

	ids := make([]uuid.UUID, 0, 10)

	for rows.Next() {
		var id uuid.UUID

		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}

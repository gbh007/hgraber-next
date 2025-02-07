package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (d *Database) FileStorages(ctx context.Context) ([]fsmodel.FileStorageSystem, error) {
	builder := squirrel.Select(model.FileStorageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("file_storages")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	out := make([]fsmodel.FileStorageSystem, 0, 10)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
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

func (d *Database) FileStorage(ctx context.Context, id uuid.UUID) (fsmodel.FileStorageSystem, error) {
	builder := squirrel.Select(model.FileStorageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("file_storages").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return fsmodel.FileStorageSystem{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	fs := fsmodel.FileStorageSystem{}

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.FileStorageScanner(&fs))
	if err != nil {
		return fsmodel.FileStorageSystem{}, fmt.Errorf("scan: %w", err)
	}

	return fs, nil
}

func (d *Database) NewFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error {
	builder := squirrel.Insert("file_storages").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"id":                   fs.ID,
			"name":                 fs.Name,
			"description":          model.StringToDB(fs.Description),
			"agent_id":             model.UUIDToDB(fs.AgentID),
			"path":                 model.StringToDB(fs.Path),
			"download_priority":    fs.DownloadPriority,
			"deduplicate_priority": fs.DeduplicatePriority,
			"highway_enabled":      fs.HighwayEnabled,
			"highway_addr":         model.URLToDB(fs.HighwayAddr),
			"created_at":           fs.CreatedAt,
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

func (d *Database) UpdateFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error {
	builder := squirrel.Update("file_storages").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"name":                 fs.Name,
			"description":          model.StringToDB(fs.Description),
			"agent_id":             model.UUIDToDB(fs.AgentID),
			"path":                 model.StringToDB(fs.Path),
			"download_priority":    fs.DownloadPriority,
			"deduplicate_priority": fs.DeduplicatePriority,
			"highway_enabled":      fs.HighwayEnabled,
			"highway_addr":         model.URLToDB(fs.HighwayAddr),
		}).
		Where(squirrel.Eq{
			"id": fs.ID,
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

func (d *Database) DeleteFileStorage(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("file_storages").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
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

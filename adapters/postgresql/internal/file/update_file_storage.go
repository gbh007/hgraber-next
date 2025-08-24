package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (repo *FileRepo) UpdateFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error {
	builder := squirrel.Update("file_storages").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
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

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

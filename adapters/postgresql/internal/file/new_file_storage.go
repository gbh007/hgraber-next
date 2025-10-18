package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (repo *FileRepo) NewFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error {
	table := model.FileStorageTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnID():                  fs.ID,
			table.ColumnName():                fs.Name,
			table.ColumnDescription():         model.StringToDB(fs.Description),
			table.ColumnAgentID():             model.UUIDToDB(fs.AgentID),
			table.ColumnPath():                model.StringToDB(fs.Path),
			table.ColumnDownloadPriority():    fs.DownloadPriority,
			table.ColumnDeduplicatePriority(): fs.DeduplicatePriority,
			table.ColumnHighwayEnabled():      fs.HighwayEnabled,
			table.ColumnHighwayAddr():         model.URLToDB(fs.HighwayAddr),
			table.ColumnCreatedAt():           fs.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

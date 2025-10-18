package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

var FileStorageTable FileStorage

type FileStorage struct{}

func (FileStorage) Name() string {
	return "file_storages"
}

func (FileStorage) ColumnID() string                  { return "id" }
func (FileStorage) ColumnName() string                { return "name" }
func (FileStorage) ColumnDescription() string         { return "description" }
func (FileStorage) ColumnAgentID() string             { return "agent_id" }
func (FileStorage) ColumnPath() string                { return "path" }
func (FileStorage) ColumnDownloadPriority() string    { return "download_priority" }
func (FileStorage) ColumnDeduplicatePriority() string { return "deduplicate_priority" }
func (FileStorage) ColumnHighwayEnabled() string      { return "highway_enabled" }
func (FileStorage) ColumnHighwayAddr() string         { return "highway_addr" }
func (FileStorage) ColumnCreatedAt() string           { return "created_at" }

func (fs FileStorage) Columns() []string {
	return []string{
		fs.ColumnID(),
		fs.ColumnName(),
		fs.ColumnDescription(),
		fs.ColumnAgentID(),
		fs.ColumnPath(),
		fs.ColumnDownloadPriority(),
		fs.ColumnDeduplicatePriority(),
		fs.ColumnHighwayEnabled(),
		fs.ColumnHighwayAddr(),
		fs.ColumnCreatedAt(),
	}
}

func (FileStorage) Scanner(fs *fsmodel.FileStorageSystem) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			description sql.NullString
			agentID     uuid.NullUUID
			path        sql.NullString
			highwayAddr sql.NullString
		)

		err := rows.Scan(
			&fs.ID,
			&fs.Name,
			&description,
			&agentID,
			&path,
			&fs.DownloadPriority,
			&fs.DeduplicatePriority,
			&fs.HighwayEnabled,
			&highwayAddr,
			&fs.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		if highwayAddr.Valid {
			fs.HighwayAddr, err = url.Parse(highwayAddr.String)
			if err != nil {
				return fmt.Errorf("convert to entity: %w", err)
			}
		}

		fs.Description = description.String
		fs.AgentID = agentID.UUID
		fs.Path = path.String

		return nil
	}
}

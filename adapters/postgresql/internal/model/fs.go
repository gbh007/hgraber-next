package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func FileStorageColumns() []string {
	return []string{
		"id",
		"name",
		"description",
		"agent_id",
		"path",
		"download_priority",
		"deduplicate_priority",
		"highway_enabled",
		"highway_addr",
		"created_at",
	}
}

func FileStorageScanner(fs *core.FileStorageSystem) RowScanner {
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

package entities

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type FileStorageSystem struct {
	ID                  uuid.UUID
	Name                string
	Description         string
	AgentID             uuid.UUID
	Path                string
	DownloadPriority    int
	DeduplicatePriority int
	HighwayEnabled      bool
	HighwayAddr         *url.URL
	CreatedAt           time.Time
}

func (fs FileStorageSystem) NotAvailable() bool {
	return fs.AgentID == uuid.Nil && fs.Path == ""
}

type FSWithStatus struct {
	Info FileStorageSystem

	// Признак того что это устаревшее хранилище через конфиг
	IsLegacy bool

	DBFile         *FSFilesInfo
	DBInvalidFile  *FSFilesInfo
	DBDetachedFile *FSFilesInfo
}

type FSFilesInfo struct {
	Count int64
	Size  int64
}

type FileTransfer struct {
	FileID uuid.UUID
	FSID   uuid.UUID
}

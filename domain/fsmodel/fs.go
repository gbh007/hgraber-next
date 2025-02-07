package fsmodel

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

const ApproximateFSCount = 10

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

	DBFile         *core.SizeWithCount
	DBInvalidFile  *core.SizeWithCount
	DBDetachedFile *core.SizeWithCount

	AvailableSize int64
}

type FileTransfer struct {
	FileID uuid.UUID
	FSID   uuid.UUID
}

type FSState struct {
	FileIDs []uuid.UUID
	Files   []FSStateFile

	TotalFileCount int64
	TotalFileSize  int64
	AvailableSize  int64
}

type FSStateFile struct {
	ID        uuid.UUID
	Size      int64
	CreatedAt time.Time
}

type FileFilter struct {
	FSID       *uuid.UUID
	BookID     *uuid.UUID
	PageNumber *int
}

package entities

import (
	"time"

	"github.com/google/uuid"
)

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

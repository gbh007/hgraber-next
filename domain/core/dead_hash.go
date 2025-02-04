package core

import (
	"time"

	"github.com/google/uuid"
)

type DeadHash struct {
	FileHash
	CreatedAt time.Time
}

type PageWithDeadHash struct {
	Page
	FSID        *uuid.UUID
	HasDeadHash bool
}

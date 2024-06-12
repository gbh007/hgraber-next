package entities

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID
	Filename  string
	Ext       string
	Md5Sum    string
	Sha256Sum string
	Size      int64
	CreateAt  time.Time
}

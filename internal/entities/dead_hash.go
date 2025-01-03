package entities

import "time"

type DeadHash struct {
	FileHash
	CreatedAt time.Time
}

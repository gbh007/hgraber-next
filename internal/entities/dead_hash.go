package entities

import "time"

type DeadHash struct {
	FileHash
	CreatedAt time.Time
}

type PageWithDeadHash struct {
	Page
	HasDeadHash bool
}

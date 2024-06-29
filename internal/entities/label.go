package entities

import (
	"time"

	"github.com/google/uuid"
)

type BookLabel struct {
	BookID     uuid.UUID
	PageNumber int
	Name       string
	Value      string
	CreateAt   time.Time
}

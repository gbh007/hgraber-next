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

type BookLabelPreset struct {
	Name        string
	Values      []string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

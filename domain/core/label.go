package core

import (
	"time"

	"github.com/google/uuid"
)

const (
	LabelNameRebuildOriginName = "rebuild:origin:name"
	LabelNameRebuildOriginID   = "rebuild:origin:id"
	LabelNameRebuildOriginURL  = "rebuild:origin:url"
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

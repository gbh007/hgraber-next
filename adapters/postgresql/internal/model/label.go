package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type BookLabel struct {
	BookID     uuid.UUID `db:"book_id"`
	PageNumber int       `db:"page_number"`
	Name       string    `db:"name"`
	Value      string    `db:"value"`
	CreateAt   time.Time `db:"create_at"`
}

func (bl BookLabel) ToEntity() (core.BookLabel, error) {
	return core.BookLabel{
		BookID:     bl.BookID,
		PageNumber: bl.PageNumber,
		Name:       bl.Name,
		Value:      bl.Value,
		CreateAt:   bl.CreateAt,
	}, nil
}

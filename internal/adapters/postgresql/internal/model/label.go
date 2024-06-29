package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type BookLabel struct {
	BookID     string        `db:"book_id"`
	PageNumber sql.NullInt32 `db:"page_number"`
	Name       string        `db:"name"`
	Value      string        `db:"value"`
	CreateAt   time.Time     `db:"create_at"`
}

func (bl BookLabel) ToEntity() (entities.BookLabel, error) {
	id, err := uuid.Parse(bl.BookID)
	if err != nil {
		return entities.BookLabel{}, err
	}

	return entities.BookLabel{
		BookID:     id,
		PageNumber: int(bl.PageNumber.Int32),
		Name:       bl.Name,
		Value:      bl.Value,
		CreateAt:   bl.CreateAt,
	}, nil
}

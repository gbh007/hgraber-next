package entities

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID               uuid.UUID
	Name             string
	OriginURL        string
	PageCount        int
	AttributesParsed bool
	CreateAt         time.Time
}

// FIXME: подумать что делать с такими моделями
type BookFull struct {
	Book
	Pages      []Page
	Attributes map[string][]string
}

type BookFilter struct {
	Limit    int
	Offset   int
	NewFirst bool
}

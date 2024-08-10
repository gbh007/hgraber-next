package entities

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID               uuid.UUID
	Name             string
	OriginURL        *url.URL
	PageCount        int
	AttributesParsed bool
	CreateAt         time.Time
}

func (b Book) IsLoaded() bool {
	return b.AttributesParsed && b.PageCount > 0 && b.Name != ""
}

// FIXME: подумать что делать с такими моделями
type BookFull struct {
	Book       Book
	Pages      []Page
	Attributes map[string][]string
	Labels     []BookLabel
}

func (b BookFull) IsLoaded() bool {
	if !b.Book.IsLoaded() {
		return false
	}

	for _, p := range b.Pages {
		if !p.IsLoaded() {
			return false
		}
	}

	return true
}

type BookWithAgent struct {
	Book
	AgentID uuid.UUID
}

type BookFullWithAgent struct {
	BookFull
	AgentID uuid.UUID
}

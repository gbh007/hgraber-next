package massloadmodel

import (
	"net/url"
	"time"
)

type Massload struct {
	ID          int
	Name        string
	Description string
	Flags       []string
	PageSize    *int64
	FileSize    *int64
	CreatedAt   time.Time
	UpdatedAt   time.Time

	ExternalLinks []ExternalLink
	Attributes    []Attribute
}

type ExternalLink struct {
	URL       url.URL
	CreatedAt time.Time
}

type Attribute struct {
	Code      string
	Value     string
	PageSize  *int64
	FileSize  *int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Flag struct {
	Code        string
	Name        string
	Description string
	CreatedAt   time.Time
}

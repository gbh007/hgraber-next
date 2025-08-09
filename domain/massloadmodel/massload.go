package massloadmodel

import (
	"net/url"
	"time"
)

type Massload struct {
	ID             int
	Description    string
	IsDeduplicated bool
	PageSize       *int64
	FileSize       *int64
	CreatedAt      time.Time
	UpdatedAt      time.Time

	ExternalLinks []MassloadExternalLink
	Attributes    []MassloadAttribute
}

type MassloadExternalLink struct {
	URL       url.URL
	CreatedAt time.Time
}

type MassloadAttribute struct {
	AttrCode  string
	AttrValue string
	PageSize  *int64
	FileSize  *int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

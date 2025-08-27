package massloadmodel

import (
	"net/url"
	"time"
)

const (
	FilterOrderByID FilterOrderBy = iota
	FilterOrderByName
	FilterOrderByPageSize
	FilterOrderByFileSize
)

const (
	FilterAttributeTypeNone FilterAttributeType = iota
	FilterAttributeTypeLike
	FilterAttributeTypeIn
)

type (
	FilterOrderBy       byte
	FilterAttributeType byte
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
	Code            string
	Name            string
	Description     string
	TextColor       string
	BackgroundColor string
	OrderWeight     int
	CreatedAt       time.Time
}

type Filter struct {
	OrderBy FilterOrderBy
	Desc    bool
	Fields  FilterFields
}

type FilterFields struct {
	Name          string
	Attributes    []FilterAttribute
	Flags         []string
	ExcludedFlags []string
	ExternalLink  string
}

type FilterAttribute struct {
	Code   string
	Type   FilterAttributeType
	Values []string
}

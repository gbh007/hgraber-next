package external

import (
	"time"
)

type Info struct {
	Version string `json:"version"`
	Meta    Meta   `json:"meta"`
	Data    Book   `json:"data"`
}

type Meta struct {
	Exported       time.Time `json:"exported"`
	ServiceVersion string    `json:"service_version,omitempty"`
	ServiceName    string    `json:"service_name,omitempty"`
}

type Book struct {
	Name             string      `json:"name"`
	OriginURL        string      `json:"origin_url,omitempty"`
	PageCount        int         `json:"page_count"`
	CreateAt         time.Time   `json:"create_at"`
	AttributesParsed bool        `json:"attributes_parsed"`
	Attributes       []Attribute `json:"attributes,omitempty"`
	Pages            []Page      `json:"pages,omitempty"`
	Labels           []Label     `json:"labels,omitempty"`
}

type Page struct {
	PageNumber int       `json:"page_number"`
	Ext        string    `json:"ext"`
	OriginURL  string    `json:"origin_url,omitempty"`
	CreateAt   time.Time `json:"create_at"`
	Downloaded bool      `json:"downloaded,omitempty"`
	LoadAt     time.Time `json:"load_at,omitempty"`
	Labels     []Label   `json:"labels,omitempty"`
}

type Label struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	CreateAt time.Time `json:"create_at"`
}

type Attribute struct {
	Code   string   `json:"code"`
	Values []string `json:"values"`
}

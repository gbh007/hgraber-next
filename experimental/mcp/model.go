package mcp

import "github.com/google/uuid"

type bookData struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	OriginURL  string    `json:"origin_url,omitempty"`
	PageCount  int       `json:"page_count"`
	Tags       []string  `json:"tags,omitempty"`
	Authors    []string  `json:"authors,omitempty"`
	Categories []string  `json:"categories,omitempty"`
	Characters []string  `json:"characters,omitempty"`
	Groups     []string  `json:"groups,omitempty"`
	Languages  []string  `json:"languages,omitempty"`
	Parodies   []string  `json:"parodies,omitempty"`
}

type attributeData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Count int    `json:"count"`
}

type hProxyValue struct {
	Value     string `json:"value"`
	OriginURL string `json:"origin_url,omitempty"`
}

type hProxyBookData struct {
	Name       string        `json:"name"`
	PageCount  int           `json:"page_count,omitempty"`
	SystemIDs  []uuid.UUID   `json:"system_ids,omitempty"`
	OriginURL  string        `json:"origin_url,omitempty"`
	Tags       []hProxyValue `json:"tags,omitempty"`
	Authors    []hProxyValue `json:"authors,omitempty"`
	Categories []hProxyValue `json:"categories,omitempty"`
	Characters []hProxyValue `json:"characters,omitempty"`
	Groups     []hProxyValue `json:"groups,omitempty"`
	Languages  []hProxyValue `json:"languages,omitempty"`
	Parodies   []hProxyValue `json:"parodies,omitempty"`
}

type hProxyListBookData struct {
	Books    []hProxyBookData `json:"books,omitempty"`
	Pages    []hProxyValue    `json:"pages,omitempty"`
	NextPage string           `json:"next_page,omitempty"`
}

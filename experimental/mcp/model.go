package mcp

import "github.com/google/uuid"

type bookData struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	PageCount  int       `json:"page_count"`
	Tags       []string  `json:"tags,omitempty"`
	Authors    []string  `json:"authors,omitempty"`
	Categories []string  `json:"categories,omitempty"`
	Characters []string  `json:"characters,omitempty"`
	Groups     []string  `json:"groups,omitempty"`
	Languages  []string  `json:"languages,omitempty"`
	Parodies   []string  `json:"parodies,omitempty"`
}

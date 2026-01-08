package mcp

import "github.com/google/uuid"

type bookData struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	PageCount int       `json:"page_count"`
	Tags      []string  `json:"tags,omitempty"`
}

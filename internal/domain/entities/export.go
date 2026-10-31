package entities

import (
	"io"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type ImportInfo struct {
	BookID     uuid.UUID
	BookURL    *url.URL
	FSPath     string
	ImportedAt time.Time
}

type ImportData struct {
	BookID   uuid.UUID
	BookName string
	BookURL  *url.URL
	Body     io.Reader
}

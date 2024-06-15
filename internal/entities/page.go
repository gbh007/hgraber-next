package entities

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Page struct {
	BookID     uuid.UUID
	PageNumber int
	Ext        string
	OriginURL  *url.URL
	CreateAt   time.Time
	Downloaded bool
	LoadAt     time.Time
	FileID     uuid.UUID
}

type PageForDownload struct {
	BookID     uuid.UUID
	PageNumber int
	Ext        string
	BookURL    *url.URL
	ImageURL   *url.URL
}

type PageForDownloadWithAgent struct {
	PageForDownload
	AgentID uuid.UUID
}

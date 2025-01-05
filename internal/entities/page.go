package entities

import (
	"net/url"
	"strconv"
	"strings"
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

func (p Page) IsLoaded() bool {
	return p.Downloaded && p.FileID != uuid.Nil
}

func (p Page) Filename() string {
	if strings.HasPrefix(p.Ext, ".") {
		return strconv.Itoa(p.PageNumber) + p.Ext
	}

	return strconv.Itoa(p.PageNumber) + "." + p.Ext
}

func (p Page) ToAgentBookDetailsPagesItem() AgentBookDetailsPagesItem {
	var u url.URL

	if p.OriginURL != nil {
		u = *p.OriginURL
	}

	return AgentBookDetailsPagesItem{
		PageNumber: p.PageNumber,
		URL:        u,
		Filename:   p.Filename(),
	}
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

type PageWithHash struct {
	Page
	FileHash
}

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
	BookID     uuid.UUID
	PageNumber int
	Ext        string
	OriginURL  *url.URL
	Downloaded bool
	FileID     uuid.UUID
	Md5Sum     string
	Sha256Sum  string
	Size       int64
}

func (p PageWithHash) Hash() FileHash {
	return FileHash{
		Md5Sum:    p.Md5Sum,
		Sha256Sum: p.Sha256Sum,
		Size:      p.Size,
	}
}

func (p PageWithHash) Page() Page {
	return Page{
		BookID:     p.BookID,
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		OriginURL:  p.OriginURL,
		Downloaded: p.Downloaded,
		FileID:     p.FileID,
	}
}

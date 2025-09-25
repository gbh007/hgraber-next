//nolint:decorder // будет исправлено позднее
package core

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	PageNumberForPreview = 1
	AvgPageCountInBook   = 30
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

type PageForDownload struct {
	BookID     uuid.UUID
	PageNumber int
	Ext        string
	BookURL    *url.URL
	ImageURL   *url.URL
}

type PageWithHash struct {
	Page
	FileHash

	FSID uuid.UUID
}

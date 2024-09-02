package entities

import (
	"io"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID            uuid.UUID
	Name          string
	Addr          url.URL
	Token         string
	CanParse      bool
	CanParseMulti bool
	CanExport     bool
	Priority      int
	CreateAt      time.Time
}

type AgentFilter struct {
	CanParse      bool
	CanParseMulti bool
	CanExport     bool
}

type AgentBookDetails struct {
	URL        url.URL
	Name       string
	PageCount  int
	Attributes []AgentBookDetailsAttributesItem
	Pages      []AgentBookDetailsPagesItem
}

type AgentBookDetailsAttributesItem struct {
	Code   string
	Values []string
}

type AgentBookDetailsPagesItem struct {
	PageNumber int
	URL        url.URL
	Filename   string
}

type AgentBookCheckResult struct {
	URL                url.URL
	IsUnsupported      bool
	IsPossible         bool
	HasError           bool
	PossibleDuplicates []url.URL
	ErrorReason        string
}

type AgentPageURL struct {
	BookURL  url.URL
	ImageURL url.URL
}

type AgentPageCheckResult struct {
	BookURL       url.URL
	ImageURL      url.URL
	IsUnsupported bool
	IsPossible    bool
	HasError      bool
	ErrorReason   string
}

type AgentExportData struct {
	BookID   uuid.UUID
	BookName string
	BookURL  *url.URL
	Body     io.Reader
}

package serveragentdto

import (
	"io"
	"net/url"

	"github.com/google/uuid"
)

type AgentBookCheckResult struct {
	URL           url.URL
	IsUnsupported bool
	IsPossible    bool
	HasError      bool
	ErrorReason   string
}

type AgentPageURL struct {
	BookURL  url.URL
	ImageURL url.URL
}

type AgentExportData struct {
	BookID   uuid.UUID
	BookName string
	BookURL  *url.URL
	Body     io.Reader
}

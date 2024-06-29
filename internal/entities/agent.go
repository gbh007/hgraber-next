package entities

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID        uuid.UUID
	Name      string
	Addr      url.URL
	Token     string
	CanParse  bool
	CanExport bool
	Priority  int
	CreateAt  time.Time
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

type AgentStatus struct {
	StartAt   time.Time
	IsOK      bool
	IsWarning bool
	IsError   bool
	Problems  []AgentStatusProblem
}

type AgentStatusProblem struct {
	IsInfo    bool
	IsWarning bool
	IsError   bool
	Details   string
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

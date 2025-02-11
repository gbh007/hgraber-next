package agentmodel

import (
	"io"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

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

func BookContainerToAgentBookDetails(b core.BookContainer) AgentBookDetails {
	var u url.URL

	if b.Book.OriginURL != nil {
		u = *b.Book.OriginURL
	}

	return AgentBookDetails{
		URL:       u,
		Name:      b.Book.Name,
		PageCount: b.Book.PageCount,
		Attributes: pkg.MapToSlice(b.Attributes, func(code string, values []string) AgentBookDetailsAttributesItem {
			return AgentBookDetailsAttributesItem{
				Code:   code,
				Values: values,
			}
		}),
		Pages: pkg.Map(b.Pages, func(p core.Page) AgentBookDetailsPagesItem {
			return PageToAgentBookDetailsPagesItem(p)
		}),
	}
}

func PageToAgentBookDetailsPagesItem(p core.Page) AgentBookDetailsPagesItem {
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

package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) ExportArchive(ctx context.Context, data serveragentdto.AgentExportData) error {
	err := a.rawClient.APIV2FsArchivePost(
		ctx,
		agentapi.APIV2FsArchivePostReq{
			Data: data.Body,
		},
		agentapi.APIV2FsArchivePostParams{
			BookID:   data.BookID,
			BookName: data.BookName,
			BookURL:  optURL(data.BookURL),
		},
	)
	if err != nil {
		return enrichError(err)
	}

	return nil
}

func optURL(u *url.URL) agentapi.OptURI {
	if u == nil {
		return agentapi.OptURI{}
	}

	return agentapi.NewOptURI(*u)
}

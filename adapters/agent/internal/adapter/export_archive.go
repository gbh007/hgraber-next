package adapter

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) ExportArchive(ctx context.Context, data agentmodel.AgentExportData) error {
	err := a.rawClient.APIImportArchivePost(
		ctx,
		agentapi.APIImportArchivePostReq{
			Data: data.Body,
		},
		agentapi.APIImportArchivePostParams{
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

package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (c *Controller) APIExportArchivePost(ctx context.Context, req agentAPI.APIExportArchivePostReq, params agentAPI.APIExportArchivePostParams) (agentAPI.APIExportArchivePostRes, error) {
	_, err := c.exportUseCases.ImportArchive(ctx, req.Data, true, false)
	if err != nil {
		return &agentAPI.APIExportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	return &agentAPI.APIExportArchivePostNoContent{}, nil
}

package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIExportArchivePost(ctx context.Context, req agentapi.APIExportArchivePostReq, params agentapi.APIExportArchivePostParams) (agentapi.APIExportArchivePostRes, error) {
	_, err := c.exportUseCases.ImportArchive(ctx, req.Data, true, false)
	if err != nil {
		return &agentapi.APIExportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIExportArchivePostNoContent{}, nil
}

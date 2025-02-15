package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIImportArchivePost(ctx context.Context, req agentapi.APIImportArchivePostReq, params agentapi.APIImportArchivePostParams) (agentapi.APIImportArchivePostRes, error) {
	_, err := c.exportUseCases.ImportArchive(ctx, req.Data, true, false)
	if err != nil {
		return &agentapi.APIImportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIImportArchivePostNoContent{}, nil
}

package apiagent

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIImportArchivePost(
	ctx context.Context,
	req agentapi.APIImportArchivePostReq,
	params agentapi.APIImportArchivePostParams,
) error {
	_, err := c.exportUseCases.ImportArchive(ctx, req.Data, true, false)
	if err != nil {
		return apiError{
			Code:      http.StatusInternalServerError,
			InnerCode: ExportUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

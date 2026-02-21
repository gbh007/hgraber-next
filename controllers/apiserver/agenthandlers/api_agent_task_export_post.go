package agenthandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentTaskExportPost(
	ctx context.Context,
	req *serverapi.APIAgentTaskExportPostReq,
) error {
	err := c.exportUseCases.Export(
		ctx,
		req.Exporter,
		apiservercore.ConvertAPIBookFilter(req.BookFilter),
		req.DeleteAfter.Value,
	)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentTaskExportPost(
	ctx context.Context,
	req *serverapi.APIAgentTaskExportPostReq,
) error {
	return c.exportUseCases.Export(
		ctx,
		req.Exporter,
		apiservercore.ConvertAPIBookFilter(req.BookFilter),
		req.DeleteAfter.Value,
	)
}

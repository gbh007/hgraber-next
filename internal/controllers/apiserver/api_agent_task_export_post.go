package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
)

func (c *Controller) APIAgentTaskExportPost(ctx context.Context, req *server.APIAgentTaskExportPostReq) (server.APIAgentTaskExportPostRes, error) {
	err := c.exportUseCases.Export(ctx, req.Exporter, req.From, req.To)
	if err != nil {
		return &server.APIAgentTaskExportPostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIAgentTaskExportPostNoContent{}, nil
}

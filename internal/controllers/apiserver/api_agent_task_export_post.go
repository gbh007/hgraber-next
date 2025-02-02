package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIAgentTaskExportPost(ctx context.Context, req *serverAPI.APIAgentTaskExportPostReq) (serverAPI.APIAgentTaskExportPostRes, error) {
	err := c.exportUseCases.Export(ctx, req.Exporter, convertAPIBookFilter(req.BookFilter), req.DeleteAfter.Value)
	if err != nil {
		return &serverAPI.APIAgentTaskExportPostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAgentTaskExportPostNoContent{}, nil
}

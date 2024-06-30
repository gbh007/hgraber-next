package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
)

func (c *Controller) APISystemImportArchivePost(ctx context.Context, req server.APISystemImportArchivePostReq) (server.APISystemImportArchivePostRes, error) {
	id, err := c.exportUseCases.ImportArchive(ctx, req.Data)
	if err != nil {
		return &server.APISystemImportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APISystemImportArchivePostOK{
		ID: id,
	}, nil
}

package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APISystemImportArchivePost(ctx context.Context, req serverapi.APISystemImportArchivePostReq) (serverapi.APISystemImportArchivePostRes, error) {
	id, err := c.exportUseCases.ImportArchive(ctx, req.Data, false, true) // FIXME: возможно все таки стоит проверять на дубли.
	if err != nil {
		return &serverapi.APISystemImportArchivePostInternalServerError{
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APISystemImportArchivePostOK{
		ID: id,
	}, nil
}

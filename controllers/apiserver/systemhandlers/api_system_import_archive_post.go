package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *SystemHandlersController) APISystemImportArchivePost(ctx context.Context, req serverAPI.APISystemImportArchivePostReq) (serverAPI.APISystemImportArchivePostRes, error) {
	id, err := c.exportUseCases.ImportArchive(ctx, req.Data, false, true) // FIXME: возможно все таки стоит проверять на дубли.
	if err != nil {
		return &serverAPI.APISystemImportArchivePostInternalServerError{
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemImportArchivePostOK{
		ID: id,
	}, nil
}

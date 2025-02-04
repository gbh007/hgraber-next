package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APISystemImportArchivePost(ctx context.Context, req serverAPI.APISystemImportArchivePostReq) (serverAPI.APISystemImportArchivePostRes, error) {
	id, err := c.exportUseCases.ImportArchive(ctx, req.Data, false, true) // FIXME: возможно все таки стоит проверять на дубли.
	if err != nil {
		return &serverAPI.APISystemImportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemImportArchivePostOK{
		ID: id,
	}, nil
}

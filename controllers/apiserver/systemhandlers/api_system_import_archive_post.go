package systemhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APISystemImportArchivePost(
	ctx context.Context,
	req serverapi.APISystemImportArchivePostReq,
) (*serverapi.APISystemImportArchivePostOK, error) {
	// FIXME: возможно все таки стоит проверять на дубли.
	id, err := c.exportUseCases.ImportArchive(
		ctx,
		req.Data,
		false,
		true,
	)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APISystemImportArchivePostOK{
		ID: id,
	}, nil
}

package fshandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsGetPost(
	ctx context.Context,
	req *serverapi.APIFsGetPostReq,
) (*serverapi.FileSystemInfo, error) {
	fs, err := c.fsUseCases.FileStorage(ctx, req.ID)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   err.Error(),
		}
	}

	result := apiservercore.ConvertFileSystemInfoToAPI(fs)

	return &result, nil
}

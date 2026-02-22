package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsGetPost(
	ctx context.Context,
	req *serverapi.APIFsGetPostReq,
) (*serverapi.FileSystemInfo, error) {
	fs, err := c.fsUseCases.FileStorage(ctx, req.ID)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	result := apiservercore.ConvertFileSystemInfoToAPI(fs)

	return &result, nil
}

package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoDeletePostReq,
) (serverapi.APIMassloadInfoDeletePostRes, error) {
	err := c.massloadUseCases.DeleteMassload(ctx, req.ID)
	if err != nil {
		return &serverapi.APIMassloadInfoDeletePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoDeletePostNoContent{}, nil
}

package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagDeletePostReq,
) (serverapi.APIMassloadFlagDeletePostRes, error) {
	err := c.massloadUseCases.DeleteMassloadFlag(ctx, req.Code)
	if err != nil {
		return &serverapi.APIMassloadFlagDeletePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadFlagDeletePostNoContent{}, nil
}

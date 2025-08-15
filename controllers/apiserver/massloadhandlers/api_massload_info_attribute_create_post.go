package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoAttributeCreatePost(ctx context.Context, req *serverapi.APIMassloadInfoAttributeCreatePostReq) (serverapi.APIMassloadInfoAttributeCreatePostRes, error) {
	err := c.massloadUseCases.CreateMassloadAttribute(ctx, req.MassloadID, req.Code, req.Value)
	if err != nil {
		return &serverapi.APIMassloadInfoAttributeCreatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoAttributeCreatePostNoContent{}, nil
}

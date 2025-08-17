package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoAttributeDeletePost(ctx context.Context, req *serverapi.APIMassloadInfoAttributeDeletePostReq) (serverapi.APIMassloadInfoAttributeDeletePostRes, error) {
	err := c.massloadUseCases.DeleteMassloadAttribute(ctx, req.MassloadID, req.Code, req.Value)
	if err != nil {
		return &serverapi.APIMassloadInfoAttributeDeletePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoAttributeDeletePostNoContent{}, nil
}

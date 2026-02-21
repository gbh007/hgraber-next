package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoAttributeCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoAttributeCreatePostReq,
) error {
	err := c.massloadUseCases.CreateMassloadAttribute(ctx, req.MassloadID, req.Code, req.Value)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

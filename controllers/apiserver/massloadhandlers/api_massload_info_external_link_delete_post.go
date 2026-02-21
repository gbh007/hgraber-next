package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkDeletePostReq,
) error {
	err := c.massloadUseCases.DeleteMassloadExternalLink(ctx, req.MassloadID, req.URL)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

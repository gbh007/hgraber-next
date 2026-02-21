package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkCreatePostReq,
) error {
	err := c.massloadUseCases.CreateMassloadExternalLink(ctx, req.MassloadID, massloadmodel.ExternalLink{
		URL:       req.URL,
		AutoCheck: req.AutoCheck.Value,
	})
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

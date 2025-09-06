package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkUpdatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkUpdatePostReq,
) (serverapi.APIMassloadInfoExternalLinkUpdatePostRes, error) {
	err := c.massloadUseCases.UpdateMassloadExternalLink(ctx, req.MassloadID, massloadmodel.ExternalLink{
		URL:       req.URL,
		AutoCheck: req.AutoCheck.Value,
	})
	if err != nil {
		return &serverapi.APIMassloadInfoExternalLinkUpdatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoExternalLinkUpdatePostNoContent{}, nil
}

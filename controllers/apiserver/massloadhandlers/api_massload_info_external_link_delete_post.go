package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkDeletePostReq,
) (serverapi.APIMassloadInfoExternalLinkDeletePostRes, error) {
	err := c.massloadUseCases.DeleteMassloadExternalLink(ctx, req.MassloadID, req.URL)
	if err != nil {
		return &serverapi.APIMassloadInfoExternalLinkDeletePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoExternalLinkDeletePostNoContent{}, nil
}

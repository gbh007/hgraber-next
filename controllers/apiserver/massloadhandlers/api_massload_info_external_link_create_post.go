package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkCreatePostReq,
) (serverapi.APIMassloadInfoExternalLinkCreatePostRes, error) {
	err := c.massloadUseCases.CreateMassloadExternalLink(ctx, req.MassloadID, req.URL)
	if err != nil {
		return &serverapi.APIMassloadInfoExternalLinkCreatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoExternalLinkCreatePostNoContent{}, nil
}

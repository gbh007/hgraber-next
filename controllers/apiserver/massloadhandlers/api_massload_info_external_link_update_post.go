package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkUpdatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkUpdatePostReq,
) error {
	return c.massloadUseCases.UpdateMassloadExternalLink(ctx, req.MassloadID, massloadmodel.ExternalLink{
		URL:       req.URL,
		AutoCheck: req.AutoCheck.Value,
	})
}

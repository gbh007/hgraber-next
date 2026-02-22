package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkCreatePostReq,
) error {
	return c.massloadUseCases.CreateMassloadExternalLink(ctx, req.MassloadID, massloadmodel.ExternalLink{
		URL:       req.URL,
		AutoCheck: req.AutoCheck.Value,
	})
}

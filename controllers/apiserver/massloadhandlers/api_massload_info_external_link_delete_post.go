package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoExternalLinkDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoExternalLinkDeletePostReq,
) error {
	return c.massloadUseCases.DeleteMassloadExternalLink(ctx, req.MassloadID, req.URL)
}

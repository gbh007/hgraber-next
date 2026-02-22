package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoAttributeDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoAttributeDeletePostReq,
) error {
	return c.massloadUseCases.DeleteMassloadAttribute(ctx, req.MassloadID, req.Code, req.Value)
}

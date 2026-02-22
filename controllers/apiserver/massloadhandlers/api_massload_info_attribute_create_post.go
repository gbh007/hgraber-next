package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoAttributeCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoAttributeCreatePostReq,
) error {
	return c.massloadUseCases.CreateMassloadAttribute(ctx, req.MassloadID, req.Code, req.Value)
}

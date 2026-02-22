package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagDeletePostReq,
) error {
	return c.massloadUseCases.DeleteMassloadFlag(ctx, req.Code)
}

package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoDeletePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoDeletePostReq,
) error {
	return c.massloadUseCases.DeleteMassload(ctx, req.ID)
}

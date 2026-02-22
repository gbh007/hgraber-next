package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoUpdatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoUpdatePostReq,
) error {
	return c.massloadUseCases.UpdateMassload(ctx, massloadmodel.Massload{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description.Value,
		Flags:       req.Flags,
	})
}

package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagUpdatePost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagUpdatePostReq,
) error {
	return c.massloadUseCases.UpdateMassloadFlag(ctx, massloadmodel.Flag{
		Code:            req.Code,
		Name:            req.Name,
		Description:     req.Description.Value,
		TextColor:       req.TextColor.Value,
		BackgroundColor: req.BackgroundColor.Value,
		OrderWeight:     req.OrderWeight,
	})
}

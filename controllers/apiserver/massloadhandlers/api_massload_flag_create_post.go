package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagCreatePostReq,
) (serverapi.APIMassloadFlagCreatePostRes, error) {
	err := c.massloadUseCases.CreateMassloadFlag(ctx, massloadmodel.Flag{
		Code:            req.Code,
		Name:            req.Name,
		Description:     req.Description.Value,
		TextColor:       req.TextColor.Value,
		BackgroundColor: req.BackgroundColor.Value,
		OrderWeight:     req.OrderWeight,
	})
	if err != nil {
		return &serverapi.APIMassloadFlagCreatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadFlagCreatePostNoContent{}, nil
}

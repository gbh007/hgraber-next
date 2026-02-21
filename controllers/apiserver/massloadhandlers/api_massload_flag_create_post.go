package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagCreatePostReq,
) error {
	err := c.massloadUseCases.CreateMassloadFlag(ctx, massloadmodel.Flag{
		Code:            req.Code,
		Name:            req.Name,
		Description:     req.Description.Value,
		TextColor:       req.TextColor.Value,
		BackgroundColor: req.BackgroundColor.Value,
		OrderWeight:     req.OrderWeight,
	})
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

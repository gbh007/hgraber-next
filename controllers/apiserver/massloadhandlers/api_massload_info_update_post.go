package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoUpdatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoUpdatePostReq,
) error {
	err := c.massloadUseCases.UpdateMassload(ctx, massloadmodel.Massload{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description.Value,
		Flags:       req.Flags,
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

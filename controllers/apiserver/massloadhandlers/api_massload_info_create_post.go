package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoCreatePostReq,
) (*serverapi.APIMassloadInfoCreatePostOK, error) {
	id, err := c.massloadUseCases.CreateMassload(ctx, massloadmodel.Massload{
		Name:        req.Name,
		Description: req.Description.Value,
		Flags:       req.Flags,
	})
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIMassloadInfoCreatePostOK{
		ID: id,
	}, nil
}

package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoCreatePost(
	ctx context.Context,
	req *serverapi.APIMassloadInfoCreatePostReq,
) (serverapi.APIMassloadInfoCreatePostRes, error) {
	id, err := c.massloadUseCases.CreateMassload(ctx, massloadmodel.Massload{
		Name:        req.Name,
		Description: req.Description.Value,
		Flags:       req.Flags,
	})
	if err != nil {
		return &serverapi.APIMassloadInfoCreatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoCreatePostOK{
		ID: id,
	}, nil
}

package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadInfoUpdatePost(ctx context.Context, req *serverapi.APIMassloadInfoUpdatePostReq) (serverapi.APIMassloadInfoUpdatePostRes, error) {
	err := c.massloadUseCases.UpdateMassload(ctx, massloadmodel.Massload{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description.Value,
		Flags:       req.Flags,
	})
	if err != nil {
		return &serverapi.APIMassloadInfoUpdatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoUpdatePostNoContent{}, nil
}

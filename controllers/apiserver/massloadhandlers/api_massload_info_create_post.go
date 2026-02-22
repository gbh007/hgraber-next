package massloadhandlers

import (
	"context"

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
		return nil, err
	}

	return &serverapi.APIMassloadInfoCreatePostOK{
		ID: id,
	}, nil
}

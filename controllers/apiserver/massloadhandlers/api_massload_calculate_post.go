package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadCalculatePost(
	ctx context.Context,
	req *serverapi.APIMassloadCalculatePostReq,
) (serverapi.APIMassloadCalculatePostRes, error) {
	if !req.ID.Set && !req.All.Value {
		return &serverapi.APIMassloadCalculatePostBadRequest{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString("invalid target"),
		}, nil
	}

	var err error

	if req.All.Value {
		err = c.massloadUseCases.CalculateMassloads(ctx, req.Force.Value)
	} else {
		err = c.massloadUseCases.CalculateMassload(ctx, req.ID.Value, req.Force.Value)
	}

	if err != nil {
		return &serverapi.APIMassloadCalculatePostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadCalculatePostNoContent{}, nil
}

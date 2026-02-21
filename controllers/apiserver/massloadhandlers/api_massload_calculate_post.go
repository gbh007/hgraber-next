package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadCalculatePost(
	ctx context.Context,
	req *serverapi.APIMassloadCalculatePostReq,
) error {
	if !req.ID.Set && !req.All.Value {
		return apiservercore.APIError{
			Code:      http.StatusBadRequest,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   "invalid target",
		}
	}

	var err error

	if req.All.Value {
		err = c.massloadUseCases.CalculateMassloads(ctx, req.Force.Value)
	} else {
		err = c.massloadUseCases.CalculateMassload(ctx, req.ID.Value, req.Force.Value)
	}

	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

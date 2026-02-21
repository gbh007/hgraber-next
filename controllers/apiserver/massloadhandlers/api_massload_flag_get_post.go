package massloadhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagGetPost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagGetPostReq,
) (*serverapi.MassloadFlag, error) {
	flag, err := c.massloadUseCases.MassloadFlag(ctx, req.Code)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   err.Error(),
		}
	}

	resp := convertMassloadFlag(flag)

	return &resp, nil
}

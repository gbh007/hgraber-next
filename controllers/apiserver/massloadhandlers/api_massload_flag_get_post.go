package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagGetPost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagGetPostReq,
) (serverapi.APIMassloadFlagGetPostRes, error) {
	flag, err := c.massloadUseCases.MassloadFlag(ctx, req.Code)
	if err != nil {
		return &serverapi.APIMassloadFlagGetPostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	resp := convertMassloadFlag(flag)

	return &resp, nil
}

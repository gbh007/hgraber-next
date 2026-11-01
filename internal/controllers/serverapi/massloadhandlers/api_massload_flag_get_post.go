package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagGetPost(
	ctx context.Context,
	req *serverapi.APIMassloadFlagGetPostReq,
) (*serverapi.MassloadFlag, error) {
	flag, err := c.massloadUseCases.MassloadFlag(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	resp := convertMassloadFlag(flag)

	return &resp, nil
}

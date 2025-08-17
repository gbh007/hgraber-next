package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *MassloadController) APIMassloadInfoListPost(ctx context.Context, req *serverapi.APIMassloadInfoListPostReq) (serverapi.APIMassloadInfoListPostRes, error) {
	mls, err := c.massloadUseCases.Massloads(ctx)
	if err != nil {
		return &serverapi.APIMassloadInfoListPostInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoListPostOK{
		Massloads: pkg.Map(mls, convertMassloadInfo),
	}, nil
}

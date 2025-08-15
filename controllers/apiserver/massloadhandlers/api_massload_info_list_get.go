package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *MassloadController) APIMassloadInfoListGet(ctx context.Context) (serverapi.APIMassloadInfoListGetRes, error) {
	mls, err := c.massloadUseCases.Massloads(ctx)
	if err != nil {
		return &serverapi.APIMassloadInfoListGetInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadInfoListGetOK{
		Massloads: pkg.Map(mls, convertMassloadInfo),
	}, nil
}

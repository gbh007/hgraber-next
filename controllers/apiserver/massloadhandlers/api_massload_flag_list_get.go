package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *MassloadController) APIMassloadFlagListGet(ctx context.Context) (serverapi.APIMassloadFlagListGetRes, error) {
	flags, err := c.massloadUseCases.MassloadFlags(ctx)
	if err != nil {
		return &serverapi.APIMassloadFlagListGetInternalServerError{
			InnerCode: apiservercore.MassloadUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIMassloadFlagListGetOK{
		Flags: pkg.Map(flags, func(flag massloadmodel.Flag) serverapi.APIMassloadFlagListGetOKFlagsItem {
			return serverapi.APIMassloadFlagListGetOKFlagsItem{
				Code:        flag.Code,
				Name:        flag.Name,
				Description: apiservercore.OptString(flag.Description),
				CreatedAt:   flag.CreatedAt,
			}
		}),
	}, nil
}

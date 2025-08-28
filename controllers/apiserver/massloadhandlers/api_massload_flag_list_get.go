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
		Flags: pkg.Map(flags, convertMassloadFlag),
	}, nil
}

func convertMassloadFlag(flag massloadmodel.Flag) serverapi.MassloadFlag {
	return serverapi.MassloadFlag{
		Code:            flag.Code,
		Name:            flag.Name,
		Description:     apiservercore.OptString(flag.Description),
		OrderWeight:     flag.OrderWeight,
		TextColor:       apiservercore.OptString(flag.TextColor),
		BackgroundColor: apiservercore.OptString(flag.BackgroundColor),
		CreatedAt:       flag.CreatedAt,
	}
}

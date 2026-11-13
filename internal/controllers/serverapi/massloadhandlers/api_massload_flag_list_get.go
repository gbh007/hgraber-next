package massloadhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *MassloadController) APIMassloadFlagListGet(ctx context.Context) (*serverapi.APIMassloadFlagListGetOK, error) {
	flags, err := c.massloadUseCases.MassloadFlags(ctx)
	if err != nil {
		return nil, err
	}

	return &serverapi.APIMassloadFlagListGetOK{
		Flags: lo.Map(flags, func(flag massloadmodel.Flag, _ int) serverapi.MassloadFlag {
			return convertMassloadFlag(flag)
		}),
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

package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIAttributeColorListGet(ctx context.Context) (serverAPI.APIAttributeColorListGetRes, error) {
	colors, err := c.webAPIUseCases.AttributeColors(ctx)
	if err != nil {
		return &serverAPI.APIAttributeColorListGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeColorListGetOK{
		Colors: pkg.Map(colors, func(raw core.AttributeColor) serverAPI.AttributeColor {
			return serverAPI.AttributeColor(raw)
		}),
	}, nil
}

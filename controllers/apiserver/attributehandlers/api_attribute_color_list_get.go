package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeColorListGet(ctx context.Context) (serverAPI.APIAttributeColorListGetRes, error) {
	colors, err := c.webAPIUseCases.AttributeColors(ctx)
	if err != nil {
		return &serverAPI.APIAttributeColorListGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeColorListGetOK{
		Colors: pkg.Map(colors, func(raw core.AttributeColor) serverAPI.AttributeColor {
			return serverAPI.AttributeColor(raw)
		}),
	}, nil
}

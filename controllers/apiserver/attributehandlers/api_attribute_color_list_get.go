package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeColorListGet(ctx context.Context) (serverapi.APIAttributeColorListGetRes, error) {
	colors, err := c.webAPIUseCases.AttributeColors(ctx)
	if err != nil {
		return &serverapi.APIAttributeColorListGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeColorListGetOK{
		Colors: pkg.Map(colors, func(raw core.AttributeColor) serverapi.AttributeColor {
			return serverapi.AttributeColor(raw)
		}),
	}, nil
}

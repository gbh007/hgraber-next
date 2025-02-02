package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
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
		Colors: pkg.Map(colors, func(raw entities.AttributeColor) serverAPI.AttributeColor {
			return serverAPI.AttributeColor(raw)
		}),
	}, nil
}

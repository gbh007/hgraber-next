package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeColorListGet(
	ctx context.Context,
) (*serverapi.APIAttributeColorListGetOK, error) {
	colors, err := c.attributeUseCases.AttributeColors(ctx)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIAttributeColorListGetOK{
		Colors: pkg.Map(colors, func(raw core.AttributeColor) serverapi.AttributeColor {
			return serverapi.AttributeColor(raw)
		}),
	}, nil
}

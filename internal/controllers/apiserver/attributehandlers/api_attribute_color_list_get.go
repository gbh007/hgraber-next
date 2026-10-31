package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorListGet(
	ctx context.Context,
) (*serverapi.APIAttributeColorListGetOK, error) {
	colors, err := c.attributeUseCases.AttributeColors(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIAttributeColorListGetOK{
		Colors: pkg.Map(colors, func(raw core.AttributeColor) serverapi.AttributeColor {
			return serverapi.AttributeColor(raw)
		}),
	}, nil
}

package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *AttributeHandlersController) APIAttributeColorListGet(
	ctx context.Context,
) (*serverapi.APIAttributeColorListGetOK, error) {
	colors, err := c.attributeUseCases.AttributeColors(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIAttributeColorListGetOK{
		Colors: lo.Map(colors, func(raw core.AttributeColor, _ int) serverapi.AttributeColor {
			return serverapi.AttributeColor(raw)
		}),
	}, nil
}

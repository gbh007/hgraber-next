package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *AttributeHandlersController) APIAttributeRemapListGet(
	ctx context.Context,
) (*serverapi.APIAttributeRemapListGetOK, error) {
	colors, err := c.attributeUseCases.AttributeRemaps(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIAttributeRemapListGetOK{
		Remaps: lo.Map(colors, func(raw core.AttributeRemap, _ int) serverapi.AttributeRemap { //nolint:gocritic,golines,lll // не понятно в чем проблема
			return apiservercore.ConvertAttributeRemapToAPI(raw)
		}),
	}, nil
}

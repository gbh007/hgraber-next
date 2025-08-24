package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeRemapListGet(
	ctx context.Context,
) (serverapi.APIAttributeRemapListGetRes, error) {
	colors, err := c.attributeUseCases.AttributeRemaps(ctx)
	if err != nil {
		return &serverapi.APIAttributeRemapListGetInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeRemapListGetOK{
		Remaps: pkg.Map(colors, func(raw core.AttributeRemap) serverapi.AttributeRemap {
			return apiservercore.ConvertAttributeRemapToAPI(raw)
		}),
	}, nil
}

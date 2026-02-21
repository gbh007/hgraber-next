package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeRemapListGet(
	ctx context.Context,
) (*serverapi.APIAttributeRemapListGetOK, error) {
	colors, err := c.attributeUseCases.AttributeRemaps(ctx)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIAttributeRemapListGetOK{
		Remaps: pkg.Map(colors, func(raw core.AttributeRemap) serverapi.AttributeRemap { //nolint:gocritic,golines,lll // не понятно в чем проблема
			return apiservercore.ConvertAttributeRemapToAPI(raw)
		}),
	}, nil
}

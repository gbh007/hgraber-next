package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorCreatePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorCreatePostReq,
) error {
	err := c.attributeUseCases.CreateAttributeColor(ctx, core.AttributeColor{
		Code:            req.Code,
		Value:           req.Value,
		TextColor:       req.TextColor,
		BackgroundColor: req.BackgroundColor,
	})
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

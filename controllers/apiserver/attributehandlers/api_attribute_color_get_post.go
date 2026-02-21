package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorGetPost(
	ctx context.Context,
	req *serverapi.APIAttributeColorGetPostReq,
) (*serverapi.AttributeColor, error) {
	color, err := c.attributeUseCases.AttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	result := serverapi.AttributeColor(color)

	return &result, nil
}

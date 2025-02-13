package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorGetPost(ctx context.Context, req *serverapi.APIAttributeColorGetPostReq) (serverapi.APIAttributeColorGetPostRes, error) {
	color, err := c.attributeUseCases.AttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return &serverapi.APIAttributeColorGetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	result := serverapi.AttributeColor(color)

	return &result, nil
}

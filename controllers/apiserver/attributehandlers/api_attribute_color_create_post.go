package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorCreatePost(ctx context.Context, req *serverapi.AttributeColor) (serverapi.APIAttributeColorCreatePostRes, error) {
	err := c.webAPIUseCases.CreateAttributeColor(ctx, core.AttributeColor(*req))
	if err != nil {
		return &serverapi.APIAttributeColorCreatePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeColorCreatePostNoContent{}, nil
}

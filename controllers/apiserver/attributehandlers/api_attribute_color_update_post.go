package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorUpdatePost(ctx context.Context, req *serverapi.AttributeColor) (serverapi.APIAttributeColorUpdatePostRes, error) {
	err := c.attributeUseCases.UpdateAttributeColor(ctx, core.AttributeColor(*req))
	if err != nil {
		return &serverapi.APIAttributeColorUpdatePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeColorUpdatePostNoContent{}, nil
}

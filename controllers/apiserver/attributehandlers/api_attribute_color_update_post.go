package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *AttributeHandlersController) APIAttributeColorUpdatePost(ctx context.Context, req *serverAPI.AttributeColor) (serverAPI.APIAttributeColorUpdatePostRes, error) {
	err := c.webAPIUseCases.UpdateAttributeColor(ctx, core.AttributeColor(*req))
	if err != nil {
		return &serverAPI.APIAttributeColorUpdatePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeColorUpdatePostNoContent{}, nil
}

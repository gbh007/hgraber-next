package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorDeletePost(ctx context.Context, req *serverapi.APIAttributeColorDeletePostReq) (serverapi.APIAttributeColorDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteAttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return &serverapi.APIAttributeColorDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeColorDeletePostNoContent{}, nil
}

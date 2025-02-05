package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *AttributeHandlersController) APIAttributeColorDeletePost(ctx context.Context, req *serverAPI.APIAttributeColorDeletePostReq) (serverAPI.APIAttributeColorDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteAttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return &serverAPI.APIAttributeColorDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeColorDeletePostNoContent{}, nil
}

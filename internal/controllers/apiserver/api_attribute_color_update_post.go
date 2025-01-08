package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAttributeColorUpdatePost(ctx context.Context, req *serverAPI.AttributeColor) (serverAPI.APIAttributeColorUpdatePostRes, error) {
	err := c.webAPIUseCases.UpdateAttributeColor(ctx, entities.AttributeColor(*req))
	if err != nil {
		return &serverAPI.APIAttributeColorUpdatePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeColorUpdatePostNoContent{}, nil
}

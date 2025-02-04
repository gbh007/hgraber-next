package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIAttributeColorCreatePost(ctx context.Context, req *serverAPI.AttributeColor) (serverAPI.APIAttributeColorCreatePostRes, error) {
	err := c.webAPIUseCases.CreateAttributeColor(ctx, entities.AttributeColor(*req))
	if err != nil {
		return &serverAPI.APIAttributeColorCreatePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeColorCreatePostNoContent{}, nil
}

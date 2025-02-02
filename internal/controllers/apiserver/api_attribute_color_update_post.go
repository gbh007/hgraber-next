package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
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

package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAttributeColorGetPost(ctx context.Context, req *serverAPI.APIAttributeColorGetPostReq) (serverAPI.APIAttributeColorGetPostRes, error) {
	color, err := c.webAPIUseCases.AttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return &serverAPI.APIAttributeColorGetPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	result := serverAPI.AttributeColor(color)

	return &result, nil
}

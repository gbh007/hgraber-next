package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APILabelPresetCreatePost(ctx context.Context, req *serverAPI.APILabelPresetCreatePostReq) (serverAPI.APILabelPresetCreatePostRes, error) {
	err := c.webAPIUseCases.CreateLabelPreset(ctx, entities.BookLabelPreset{
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
	if err != nil {
		return &serverAPI.APILabelPresetCreatePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetCreatePostNoContent{}, nil
}

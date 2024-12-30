package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APILabelPresetUpdatePost(ctx context.Context, req *serverAPI.APILabelPresetUpdatePostReq) (serverAPI.APILabelPresetUpdatePostRes, error) {
	err := c.webAPIUseCases.UpdateLabelPreset(ctx, entities.BookLabelPreset{
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
	if err != nil {
		return &serverAPI.APILabelPresetUpdatePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetUpdatePostNoContent{}, nil
}

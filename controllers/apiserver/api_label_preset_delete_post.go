package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APILabelPresetDeletePost(ctx context.Context, req *serverAPI.APILabelPresetDeletePostReq) (serverAPI.APILabelPresetDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteLabelPreset(ctx, req.Name)
	if err != nil {
		return &serverAPI.APILabelPresetDeletePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetDeletePostNoContent{}, nil
}

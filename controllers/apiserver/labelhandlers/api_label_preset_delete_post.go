package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *LabelHandlersController) APILabelPresetDeletePost(ctx context.Context, req *serverAPI.APILabelPresetDeletePostReq) (serverAPI.APILabelPresetDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteLabelPreset(ctx, req.Name)
	if err != nil {
		return &serverAPI.APILabelPresetDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetDeletePostNoContent{}, nil
}

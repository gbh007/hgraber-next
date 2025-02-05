package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APILabelPresetUpdatePost(ctx context.Context, req *serverAPI.APILabelPresetUpdatePostReq) (serverAPI.APILabelPresetUpdatePostRes, error) {
	err := c.webAPIUseCases.UpdateLabelPreset(ctx, core.BookLabelPreset{
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
	if err != nil {
		return &serverAPI.APILabelPresetUpdatePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetUpdatePostNoContent{}, nil
}

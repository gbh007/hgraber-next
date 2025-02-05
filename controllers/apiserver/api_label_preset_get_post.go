package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APILabelPresetGetPost(ctx context.Context, req *serverAPI.APILabelPresetGetPostReq) (serverAPI.APILabelPresetGetPostRes, error) {
	raw, err := c.webAPIUseCases.LabelPreset(ctx, req.Name)
	if err != nil {
		return &serverAPI.APILabelPresetGetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetGetPostOK{
		Name:        raw.Name,
		Description: apiservercore.OptString(raw.Description),
		Values:      raw.Values,
		CreatedAt:   raw.CreatedAt,
		UpdatedAt:   apiservercore.OptTime(raw.UpdatedAt),
	}, nil
}

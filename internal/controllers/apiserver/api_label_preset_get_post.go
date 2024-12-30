package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APILabelPresetGetPost(ctx context.Context, req *serverAPI.APILabelPresetGetPostReq) (serverAPI.APILabelPresetGetPostRes, error) {
	raw, err := c.webAPIUseCases.LabelPreset(ctx, req.Name)
	if err != nil {
		return &serverAPI.APILabelPresetGetPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetGetPostOK{
		Name:        raw.Name,
		Description: optString(raw.Description),
		Values:      raw.Values,
		CreatedAt:   raw.CreatedAt,
		UpdatedAt:   optTime(raw.UpdatedAt),
	}, nil
}

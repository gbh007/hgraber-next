package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetDeletePost(
	ctx context.Context,
	req *serverapi.APILabelPresetDeletePostReq,
) (serverapi.APILabelPresetDeletePostRes, error) {
	err := c.labelUseCases.DeleteLabelPreset(ctx, req.Name)
	if err != nil {
		return &serverapi.APILabelPresetDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APILabelPresetDeletePostNoContent{}, nil
}

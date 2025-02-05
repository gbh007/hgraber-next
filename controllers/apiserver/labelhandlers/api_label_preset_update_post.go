package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetUpdatePost(ctx context.Context, req *serverapi.APILabelPresetUpdatePostReq) (serverapi.APILabelPresetUpdatePostRes, error) {
	err := c.webAPIUseCases.UpdateLabelPreset(ctx, core.BookLabelPreset{
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
	if err != nil {
		return &serverapi.APILabelPresetUpdatePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APILabelPresetUpdatePostNoContent{}, nil
}

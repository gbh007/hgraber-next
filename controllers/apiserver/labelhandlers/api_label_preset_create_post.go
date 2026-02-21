package labelhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetCreatePost(
	ctx context.Context,
	req *serverapi.APILabelPresetCreatePostReq,
) error {
	err := c.labelUseCases.CreateLabelPreset(ctx, core.BookLabelPreset{
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.LabelUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

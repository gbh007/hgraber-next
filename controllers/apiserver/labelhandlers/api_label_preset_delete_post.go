package labelhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetDeletePost(
	ctx context.Context,
	req *serverapi.APILabelPresetDeletePostReq,
) error {
	err := c.labelUseCases.DeleteLabelPreset(ctx, req.Name)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.LabelUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

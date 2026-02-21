package labelhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelSetPost(
	ctx context.Context,
	req *serverapi.APILabelSetPostReq,
) error {
	err := c.labelUseCases.SetLabel(ctx, core.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
		Value:      req.Value,
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

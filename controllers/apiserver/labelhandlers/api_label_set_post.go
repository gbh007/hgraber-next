package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelSetPost(ctx context.Context, req *serverapi.APILabelSetPostReq) (serverapi.APILabelSetPostRes, error) {
	err := c.webAPIUseCases.SetLabel(ctx, core.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
		Value:      req.Value,
	})
	if err != nil {
		return &serverapi.APILabelSetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APILabelSetPostNoContent{}, nil
}

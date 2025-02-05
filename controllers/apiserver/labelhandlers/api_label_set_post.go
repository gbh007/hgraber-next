package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *LabelHandlersController) APILabelSetPost(ctx context.Context, req *serverAPI.APILabelSetPostReq) (serverAPI.APILabelSetPostRes, error) {
	err := c.webAPIUseCases.SetLabel(ctx, core.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
		Value:      req.Value,
	})
	if err != nil {
		return &serverAPI.APILabelSetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelSetPostNoContent{}, nil
}

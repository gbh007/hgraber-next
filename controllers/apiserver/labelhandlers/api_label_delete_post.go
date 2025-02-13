package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelDeletePost(ctx context.Context, req *serverapi.APILabelDeletePostReq) (serverapi.APILabelDeletePostRes, error) {
	err := c.labelUseCases.DeleteLabel(ctx, core.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
	})
	if err != nil {
		return &serverapi.APILabelDeletePostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APILabelDeletePostNoContent{}, nil
}

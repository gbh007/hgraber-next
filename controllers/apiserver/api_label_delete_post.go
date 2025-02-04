package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APILabelDeletePost(ctx context.Context, req *serverAPI.APILabelDeletePostReq) (serverAPI.APILabelDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteLabel(ctx, core.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
	})
	if err != nil {
		return &serverAPI.APILabelDeletePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelDeletePostNoContent{}, nil
}

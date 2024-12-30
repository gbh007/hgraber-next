package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APILabelDeletePost(ctx context.Context, req *serverAPI.APILabelDeletePostReq) (serverAPI.APILabelDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteLabel(ctx, entities.BookLabel{
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

package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APILabelSetPost(ctx context.Context, req *serverAPI.APILabelSetPostReq) (serverAPI.APILabelSetPostRes, error) {
	err := c.webAPIUseCases.SetLabel(ctx, entities.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
		Value:      req.Value,
	})
	if err != nil {
		return &serverAPI.APILabelSetPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelSetPostNoContent{}, nil
}

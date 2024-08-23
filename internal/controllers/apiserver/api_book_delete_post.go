package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookDeletePost(ctx context.Context, req *serverAPI.APIBookDeletePostReq) (serverAPI.APIBookDeletePostRes, error) {
	err := c.webAPIUseCases.DeleteBook(ctx, req.ID)

	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookDeletePostNotFound{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookDeletePostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookDeletePostNoContent{}, nil
}

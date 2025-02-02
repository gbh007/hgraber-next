package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
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

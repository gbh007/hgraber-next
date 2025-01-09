package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookVerifyPost(ctx context.Context, req *serverAPI.APIBookVerifyPostReq) (serverAPI.APIBookVerifyPostRes, error) {
	err := c.webAPIUseCases.VerifyBook(ctx, req.ID, req.VerifyStatus)

	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookVerifyPostNotFound{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookVerifyPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookVerifyPostNoContent{}, nil
}

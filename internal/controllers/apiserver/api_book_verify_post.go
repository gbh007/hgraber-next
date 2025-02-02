package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
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

package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *BookHandlersController) APIBookVerifyPost(ctx context.Context, req *serverAPI.APIBookVerifyPostReq) (serverAPI.APIBookVerifyPostRes, error) {
	err := c.webAPIUseCases.VerifyBook(ctx, req.ID, req.VerifyStatus)

	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookVerifyPostNotFound{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookVerifyPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookVerifyPostNoContent{}, nil
}

package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookVerifyPost(ctx context.Context, req *serverapi.APIBookVerifyPostReq) (serverapi.APIBookVerifyPostRes, error) {
	err := c.webAPIUseCases.VerifyBook(ctx, req.ID, req.VerifyStatus)

	if errors.Is(err, core.BookNotFoundError) {
		return &serverapi.APIBookVerifyPostNotFound{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookVerifyPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookVerifyPostNoContent{}, nil
}

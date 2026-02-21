package bookhandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookStatusSetPost(
	ctx context.Context,
	req *serverapi.APIBookStatusSetPostReq,
) error {
	var err error

	switch req.Status {
	case serverapi.APIBookStatusSetPostReqStatusRebuild:
		err = c.bookUseCases.SetBookRebuild(ctx, req.ID, req.Value)
	case serverapi.APIBookStatusSetPostReqStatusVerify:
		err = c.bookUseCases.VerifyBook(ctx, req.ID, req.Value)
	}

	if errors.Is(err, core.ErrBookNotFound) {
		return apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}

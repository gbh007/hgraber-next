package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookStatusSetPost(
	ctx context.Context,
	req *serverapi.APIBookStatusSetPostReq,
) (serverapi.APIBookStatusSetPostRes, error) {
	var err error

	switch req.Status {
	case serverapi.APIBookStatusSetPostReqStatusRebuild:
		err = c.bookUseCases.SetBookRebuild(ctx, req.ID, req.Value)
	case serverapi.APIBookStatusSetPostReqStatusVerify:
		err = c.bookUseCases.VerifyBook(ctx, req.ID, req.Value)
	}

	if errors.Is(err, core.ErrBookNotFound) {
		return &serverapi.APIBookStatusSetPostNotFound{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookStatusSetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookStatusSetPostNoContent{}, nil
}

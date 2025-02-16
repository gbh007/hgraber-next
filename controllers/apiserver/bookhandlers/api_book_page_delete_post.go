package bookhandlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookPageDeletePost(ctx context.Context, req *serverapi.APIBookPageDeletePostReq) (serverapi.APIBookPageDeletePostRes, error) {
	var (
		err error
		uc  string
	)

	switch req.Type {
	case serverapi.APIBookPageDeletePostReqTypeOne:
		err = c.bookUseCases.DeletePage(ctx, req.BookID, req.PageNumber)
		uc = apiservercore.BookUseCaseCode

	case serverapi.APIBookPageDeletePostReqTypeAllCopy:
		err = c.deduplicateUseCases.DeleteAllPageByHash(ctx, req.BookID, req.PageNumber, req.SetDeadHash.Value)
		uc = apiservercore.DeduplicateUseCaseCode

	default:
		err = fmt.Errorf("unsupported type: %v", req.Type)
		uc = apiservercore.ValidationCode
	}

	if errors.Is(err, core.BookNotFoundError) ||
		errors.Is(err, core.PageNotFoundError) {
		return &serverapi.APIBookPageDeletePostNotFound{
			InnerCode: uc,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookPageDeletePostInternalServerError{
			InnerCode: uc,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookPageDeletePostNoContent{}, nil
}
